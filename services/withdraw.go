package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"time"
)

type withdraw struct {
	repo datasource.Datasources
}

func NewWithdraw(d datasource.Datasources) *withdraw {
	pl := &withdraw{d}
	return pl
}

func (pl *withdraw) Execute(cmd *schemas.Command) (err error) {
	if cmd == nil {
		err = utils.ErrorCommand
		return
	}

	cmd.ExecutedDate = time.Now()

	// get user
	user, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.From)
	if err != nil {
		return
	}

	if user.Balance < cmd.Arguments.Amount {
		err = utils.SetErrorInsufficient(user.Currency, user.Balance)
		return
	}

	balance := user.Balance - cmd.Arguments.Amount
	err = pl.repo.UpdateUserBalance(user.Id, balance)
	if err != nil {
		return
	}

	//re login to update logged user
	err = pl.repo.Login(user.Id)

	return
}
