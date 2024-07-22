package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"time"
)

type withdraw struct {
	repo ServiceDatasources
}

func NewWithdraw(d ServiceDatasources) *withdraw {
	pl := &withdraw{d}
	return pl
}

func (pl *withdraw) Execute(cmd *schemas.Command) (err error) {
	if cmd == nil {
		err = utils.ErrorCommand
		return
	}

	timeNow := time.Now()
	cmd.ExecutedDate = timeNow

	// get user
	user, err := pl.repo.UserDatasource.GetUserByAccountNumber(cmd.Arguments.AtmMachineArg.From)
	if err != nil {
		return
	}

	if user.Balance < cmd.Arguments.AtmMachineArg.Amount {
		err = utils.SetErrorInsufficient(user.Currency, user.Balance)
		return
	}

	balance := user.Balance - cmd.Arguments.AtmMachineArg.Amount
	err = pl.repo.UserDatasource.UpdateUserBalance(user.Id, balance)
	if err != nil {
		return
	}

	// insert transaction for from user
	transaction := datasource.Transaction{
		AccountNumber:   user.AccountNumber,
		InitialBalance:  user.Balance,
		Amount:          cmd.Arguments.AtmMachineArg.Amount,
		Type:            utils.Withdraw,
		TransactionDate: timeNow,
		CreditOrDebit:   utils.Credit,
	}

	_ = pl.repo.TransactionDatasource.InsertTransactionHistory(transaction)

	//re login to update logged user
	err = pl.repo.UserDatasource.Login(user.Id)

	return
}
