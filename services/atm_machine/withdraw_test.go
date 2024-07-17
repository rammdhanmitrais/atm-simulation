package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/datasource/mock"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_withdraw_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDatasource := mock.NewMockUserDatasources(ctrl)

	type args struct {
		cmd *schemas.Command
	}

	arg := args{
		cmd: &schemas.Command{
			Arguments: schemas.Arguments{
				AtmMachineArg: schemas.AtmMachineArguments{
					Amount:          500,
					ReferenceNumber: "012345",
				},
			},
		},
	}

	userData := datasource.User{
		Balance:  200,
		Currency: "$",
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
			name: "Should return error when user not found",
			args: arg,
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&datasource.User{}, utils.ErrorInvalidAccount),
			},
			wantErr: true,
		},
		{
			name: "Should return error when user balance less than amount",
			args: arg,
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&userData, nil),
			},
			wantErr: true,
		},
		{
			name: "Should return error when failed update user balance",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						AtmMachineArg: schemas.AtmMachineArguments{
							Amount:          100,
							ReferenceNumber: "012345",
						},
					},
				},
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&userData, nil),
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
						AtmMachineArg: schemas.AtmMachineArguments{
							Amount:          100,
							ReferenceNumber: "012345",
						},
					},
				},
			},
			mocks: []*gomock.Call{
				mockDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&userData, nil),
				mockDatasource.EXPECT().UpdateUserBalance(gomock.Any(), gomock.Any()).
					Times(1).
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
			pl := &withdraw{
				repo: mockDatasource,
			}
			if err := pl.Execute(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("withdraw.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
