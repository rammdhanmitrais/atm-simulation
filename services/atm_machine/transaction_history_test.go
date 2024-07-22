package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/datasource/mock"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

func Test_transactionHistory_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserDatasource := mock.NewMockUserDatasources(ctrl)
	mockTransactionDatasource := mock.NewMockTransactionDatasources(ctrl)
	userData := datasource.User{
		Balance:       200,
		Currency:      "$",
		AccountNumber: "112233",
	}

	transaction := datasource.Transaction{
		Id:              0,
		AccountNumber:   userData.AccountNumber,
		InitialBalance:  100,
		Amount:          50,
		Type:            utils.Withdraw,
		TransactionDate: time.Now(),
		CreditOrDebit:   utils.Credit,
	}

	transactions := []datasource.Transaction{transaction}
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
			name: "Should return error when logged user is nil",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						AtmMachineArg: schemas.AtmMachineArguments{
							Amount: 1500,
						},
					},
				},
			},
			mocks: []*gomock.Call{
				mockUserDatasource.EXPECT().GetLoggedUser().
					Times(1).
					Return(nil, nil),
			},
			wantErr: true,
		},
		{
			name: "Should return success",
			args: args{
				cmd: &schemas.Command{
					Arguments: schemas.Arguments{
						AtmMachineArg: schemas.AtmMachineArguments{
							Amount: 1500,
						},
					},
				},
			},
			mocks: []*gomock.Call{
				mockUserDatasource.EXPECT().GetLoggedUser().
					Times(1).
					Return(&userData, nil),
				mockTransactionDatasource.EXPECT().GetTransactionHistoriesByAccountNumber(userData.AccountNumber).
					Times(1).
					Return(transactions, nil),
				mockUserDatasource.EXPECT().UpdateLoggedUserTransactionHistory(transactions).
					Times(1),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pl := &transactionHistory{
				repo: ServiceDatasources{
					UserDatasource:        mockUserDatasource,
					TransactionDatasource: mockTransactionDatasource,
				},
			}
			if err := pl.Execute(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("transactionHistory.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
