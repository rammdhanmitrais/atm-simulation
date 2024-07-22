package atm_machine

import (
	"atm-simulation/schemas"
	"atm-simulation/utils"
)

type transactionHistory struct {
	repo ServiceDatasources
}

func NewTransactionHistory(d ServiceDatasources) *transactionHistory {
	pl := &transactionHistory{d}
	return pl
}

func (pl *transactionHistory) Execute(cmd *schemas.Command) (err error) {
	if cmd == nil {
		err = utils.ErrorCommand
		return
	}

	loggedUser, _ := pl.repo.UserDatasource.GetLoggedUser()
	if loggedUser == nil {
		err = utils.ErrorInvalidAccount

		cmd.Command = utils.LoginCommand
		return
	}

	transactions, _ := pl.repo.TransactionDatasource.GetTransactionHistoriesByAccountNumber(loggedUser.AccountNumber)
	pl.repo.UserDatasource.UpdateLoggedUserTransactionHistory(transactions)
	return
}
