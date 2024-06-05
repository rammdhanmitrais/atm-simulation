package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"fmt"
	"time"
)

type withdraw struct {
	*services
}

func NewWithdraw(s *services) *withdraw {
	pl := &withdraw{s}
	return pl
}

func (pl *withdraw) Execute(cmd *schemas.Command) (err error) {
	cmd.ExecutedDate = time.Now()

	// get user
	user, err := datasource.GetUserByAccountNumber(cmd.Arguments.From)
	if err != nil {
		return
	}

	if user.Balance < cmd.Arguments.Amount {
		err = fmt.Errorf("insufficient balance %s%d", user.Currency, user.Balance)
		return
	}

	datasource.LoggedUser.Balance = user.Balance - cmd.Arguments.Amount
	datasource.UserAccounts[user.Index] = *datasource.LoggedUser

	return
}
