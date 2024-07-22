package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"log"
	"time"
)

type fundTransfer struct {
	repo ServiceDatasources
}

func NewFundTransfer(d ServiceDatasources) *fundTransfer {
	pl := &fundTransfer{d}
	return pl
}

func (pl *fundTransfer) Execute(cmd *schemas.Command) (err error) {
	if cmd == nil {
		err = utils.ErrorCommand
		return
	}

	timeNow := time.Now()
	cmd.ExecutedDate = timeNow

	// validate argument
	if cmd.Arguments.AtmMachineArg.Amount > 1000 {
		err = utils.ErrorMaximumAmountTransfer
		return
	}

	if cmd.Arguments.AtmMachineArg.Amount < 1 {
		err = utils.ErrorMinimumAmountTransfer
		return
	}

	if cmd.Arguments.AtmMachineArg.ReferenceNumber == "" {
		err = utils.ErrorReferenceNumber
		return
	}

	// get from user
	userFrom, err := pl.repo.UserDatasource.GetUserByAccountNumber(cmd.Arguments.AtmMachineArg.From)

	if err != nil {
		return
	}

	if userFrom.Balance < cmd.Arguments.AtmMachineArg.Amount {
		err = utils.SetErrorInsufficient(userFrom.Currency, userFrom.Balance)
		return
	}

	// get from to
	userTo, err := pl.repo.UserDatasource.GetUserByAccountNumber(cmd.Arguments.AtmMachineArg.To)

	if err != nil {
		return
	}

	// modify balance from user
	fromBalance := userFrom.Balance - cmd.Arguments.AtmMachineArg.Amount
	err = pl.repo.UserDatasource.UpdateUserBalance(userFrom.Id, fromBalance)
	if err != nil {
		return
	}

	// modify balance to user
	toBalance := userTo.Balance + cmd.Arguments.AtmMachineArg.Amount
	err = pl.repo.UserDatasource.UpdateUserBalance(userTo.Id, toBalance)
	if err != nil {
		return
	}

	// insert transaction for from user
	transaction := datasource.Transaction{
		AccountNumber:       userFrom.AccountNumber,
		FromToAccountNumber: userTo.AccountNumber,
		InitialBalance:      userFrom.Balance,
		Amount:              cmd.Arguments.AtmMachineArg.Amount,
		Type:                utils.FundTransfer,
		TransactionDate:     timeNow,
		CreditOrDebit:       utils.Credit,
	}

	_ = pl.repo.TransactionDatasource.InsertTransactionHistory(transaction)

	// insert transaction for to user
	toTransaction := datasource.Transaction{
		AccountNumber:       userTo.AccountNumber,
		FromToAccountNumber: userFrom.AccountNumber,
		InitialBalance:      userTo.Balance,
		Amount:              cmd.Arguments.AtmMachineArg.Amount,
		Type:                utils.FundTransfer,
		TransactionDate:     timeNow,
		CreditOrDebit:       utils.Debit,
	}

	_ = pl.repo.TransactionDatasource.InsertTransactionHistory(toTransaction)

	//re login to update logged user
	err = pl.repo.UserDatasource.Login(userFrom.Id)
	if err != nil {
		log.Println(err)
	}

	return
}
