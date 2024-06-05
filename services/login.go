package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"fmt"
	"time"
)

type login struct {
	*services
}

func NewLogin(s *services) *login {
	pl := &login{s}
	return pl
}

func (pl *login) Execute(cmd *schemas.Command) (err error) {
	cmd.ExecutedDate = time.Now()

	// get user
	user, err := datasource.GetUserByAccountNumber(cmd.Arguments.From)

	if err != nil {
		return
	}

	if user.AccountNumber == "" || user.Pin != cmd.Arguments.Pin {
		err = fmt.Errorf("invalid account number/pin") 
		return
	}

	usr := datasource.User{
		AccountNumber: user.AccountNumber,
		Currency: user.Currency,
		Name: user.Name,
		Balance: user.Balance,
		Pin: user.Pin,
	}

	datasource.LoggedUser = &usr

	return
}
