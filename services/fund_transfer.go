package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"errors"
	"fmt"
	"time"
)

type fundTransfer struct {
	*services
}

func NewFundTransfer(s *services) *fundTransfer {
	pl := &fundTransfer{s}
	return pl
}

func (pl *fundTransfer) Execute(cmd *schemas.Command) (err error) {
	cmd.ExecutedDate = time.Now()

	// validate argument
	if cmd.Arguments.Amount > 1000 {
		err = errors.New("maximum amount to transfer is $1000")
		return
	}

	if cmd.Arguments.Amount < 1 {
		err = errors.New("minimum amount to transfer is $1")
		return
	}

	if cmd.Arguments.ReferenceNumber == "" {
		err = errors.New("invalid reference number")
		return
	}

	// get from user
	userFrom, err := datasource.GetUserByAccountNumber(cmd.Arguments.From)

	if err != nil {
		return
	}

	if userFrom.Balance < cmd.Arguments.Amount {
		err = fmt.Errorf("insufficient balance %s%d", userFrom.Currency, userFrom.Balance)
		return
	}

	// get from to
	userTo, err := datasource.GetUserByAccountNumber(cmd.Arguments.To)

	if err != nil {
		return
	}

	// modify balance from user
	datasource.LoggedUser.Balance = userFrom.Balance - cmd.Arguments.Amount
	datasource.UserAccounts[userFrom.Index] = *datasource.LoggedUser

	// modify balance to user
	datasource.UserAccounts[userTo.Index].Balance = userTo.Balance + cmd.Arguments.Amount

	return
}