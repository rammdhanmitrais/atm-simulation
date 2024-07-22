package datasource

import (
	"atm-simulation/utils"
	"reflect"
	"testing"
	"time"
)

var trxDs = NewTransactionDatasource()

func Test_transactionDatasource_GetTransactionHistoriesByAccountNumber(t *testing.T) {
	transactionHistories = append(transactionHistories, Transaction{
		AccountNumber:   "112244",
		InitialBalance:  100,
		Amount:          50,
		Type:            utils.Withdraw,
		TransactionDate: time.Now(),
		CreditOrDebit:   utils.Credit,
	})

	type args struct {
		accountNumber string
	}
	tests := []struct {
		name       string
		d          transactionDatasource
		args       args
		wantResult []Transaction
		wantErr    bool
	}{
		{
			name: "Should return success with empty transaction when user has no transaction history",
			args: args{
				accountNumber: "112233",
			},
			d:       *trxDs,
			wantErr: false,
		},
		{
			name: "Should return success return the transaction history",
			args: args{
				accountNumber: "112244",
			},
			d:          *trxDs,
			wantErr:    false,
			wantResult: transactionHistories,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := transactionDatasource{}
			gotResult, err := d.GetTransactionHistoriesByAccountNumber(tt.args.accountNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionDatasource.GetTransactionHistoriesByAccountNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("transactionDatasource.GetTransactionHistoriesByAccountNumber() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}

	transactionHistories = []Transaction{}
}

func Test_transactionDatasource_InsertTransactionHistory(t *testing.T) {
	d := *trxDs

	accountNumber1 := "112244"
	accountNumber2 := "112255"

	timeNow := time.Now()
	transaction1 := Transaction{
		Id:              0,
		AccountNumber:   accountNumber1,
		InitialBalance:  100,
		Amount:          50,
		Type:            utils.Withdraw,
		TransactionDate: timeNow,
		CreditOrDebit:   utils.Credit,
	}

	transaction2 := Transaction{
		Id:              1,
		AccountNumber:   accountNumber2,
		InitialBalance:  100,
		Amount:          20,
		Type:            utils.Withdraw,
		TransactionDate: timeNow,
		CreditOrDebit:   utils.Credit,
	}

	type args struct {
		transaction Transaction
	}
	tests := []struct {
		name          string
		d             transactionDatasource
		args          args
		wantErr       bool
		getData       Transaction
		accountNumber string
	}{
		{
			name: "Should return success and the id on transaction history is 1",
			args: args{
				transaction: transaction1,
			},
			d:             d,
			wantErr:       false,
			getData:       transaction1,
			accountNumber: accountNumber1,
		},
		{
			name: "Should return success and the id on transaction history is 2",
			args: args{
				transaction: transaction2,
			},
			d:             d,
			wantErr:       false,
			getData:       transaction2,
			accountNumber: accountNumber2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := transactionDatasource{}
			if err := d.InsertTransactionHistory(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("transactionDatasource.InsertTransactionHistory() error = %v, wantErr %v", err, tt.wantErr)
			}

			expectedData, _ := d.GetTransactionHistoriesByAccountNumber(tt.accountNumber)

			if !reflect.DeepEqual(expectedData[0], tt.getData) {
				t.Errorf("transactionDatasource.InsertTransactionHistory() = %v, want %v", tt.getData, expectedData[0])
			}
		})
	}
}
