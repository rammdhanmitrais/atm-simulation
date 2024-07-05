package services

import (
	"atm-simulation/datasource"
	"atm-simulation/datasource/mock"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_fundTransfer_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDatasource := mock.NewMockDatasources(ctrl)

	userData := datasource.User{
		Balance:  200,
		Currency: "$",
	}
	type args struct {
		cmd *schemas.Command
	}
	tests := []struct {
		name    string
		mocks   []*gomock.Call
		args    args
		wantErr bool
	}{
		{
			name:    "Should return error when command is nil",
			args:    args{},
			wantErr: true,
		},
		{
			name: "Should return error when amount greater than 1000",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount: 1500,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when amount less than 1",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount: 0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when reference number is empty",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount: 500,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Should return error when account number is not found",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount:          500,
						ReferenceNumber: "012345",
					},
				},
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(datasource.User{}, utils.ErrorInvalidAccount),
			},
			wantErr: true,
		},
		{
			name: "Should return error when user balance less than amount",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount:          500,
						ReferenceNumber: "012345",
					},
				},
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(userData, nil),
			},
			wantErr: true,
		},
		{
			name: "Should return error when failed to update the user balance",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount:          150,
						ReferenceNumber: "012345",
					},
				},
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(2).
					Return(userData, nil),

				mockDatasource.EXPECT().UpdateUserBalance(gomock.Any(), gomock.Any()).
					Times(1).
					Return(utils.ErrorInvalidAmount),
			},
			wantErr: true,
		},
		{
			name: "Should return success",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						Amount:          150,
						ReferenceNumber: "012345",
					},
				},
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(2).
					Return(userData, nil),

				mockDatasource.EXPECT().UpdateUserBalance(gomock.Any(), gomock.Any()).
					Times(2).
					Return(nil),

				mockDatasource.EXPECT().Login(gomock.Any()).
					Times(1).
					Return(nil),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &fundTransfer{
				repo: mockDatasource,
			}

			err := pl.Execute(tt.args.cmd)

			if (err != nil) != tt.wantErr {
				t.Errorf("fundTransfer.Execute() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
