package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/datasource/mock"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_login_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserDatasource := mock.NewMockUserDatasources(ctrl)
	mockTransactionDatasource := mock.NewMockTransactionDatasources(ctrl)

	type args struct {
		cmd *schemas.Command
	}

	arg := args{
		cmd: &schemas.Command{
			Arguments: schemas.Arguments{
				AtmMachineArg: schemas.AtmMachineArguments{
					Pin: "012345",
				},
			},
		},
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
				mockUserDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&datasource.User{}, utils.ErrorInvalidAccount),
			},
			wantErr: true,
		},
		{
			name: "Should return error when account number or PIN invalid",
			args: arg,
			mocks: []*gomock.Call{
				mockUserDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&datasource.User{
						Pin: "123456",
					}, nil),
			},
			wantErr: true,
		},
		{
			name: "Should return error when failed login",
			args: arg,
			mocks: []*gomock.Call{
				mockUserDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&datasource.User{
						Pin:           "012345",
						AccountNumber: "12345678",
					}, nil),
				mockUserDatasource.EXPECT().Login(gomock.Any()).
					Times(1).
					Return(utils.ErrorInvalidAccount),
			},
			wantErr: true,
		},
		{
			name: "Should return success",
			args: arg,
			mocks: []*gomock.Call{
				mockUserDatasource.EXPECT().GetUserByAccountNumber(gomock.Any()).
					Times(1).
					Return(&datasource.User{
						Pin:           "012345",
						AccountNumber: "12345678",
					}, nil),
				mockUserDatasource.EXPECT().Login(gomock.Any()).
					Times(1).
					Return(nil),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &login{
				repo: ServiceDatasources{
					mockUserDatasource,
					mockTransactionDatasource,
				},
			}
			if err := pl.Execute(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("login.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
