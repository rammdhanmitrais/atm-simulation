package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"fmt"
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
	cmd.ExecutedDate = time.Now()

	// get user
	user, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.From)
	if err != nil {
		return
	}

	if user.Balance < cmd.Arguments.Amount {
		err = fmt.Errorf("insufficient balance %s%d", user.Currency, user.Balance)
		return
	}

	balance := user.Balance - cmd.Arguments.Amount
	err = pl.repo.UpdateUserBalance(user.Id, balance)

	return
}
