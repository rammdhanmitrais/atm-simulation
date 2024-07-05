package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"time"
)

type fundTransfer struct {
	repo datasource.Datasources
}

func NewFundTransfer(d datasource.Datasources) *fundTransfer {
	pl := &fundTransfer{d}
	return pl
}

func (pl *fundTransfer) Execute(cmd *schemas.Command) (err error) {
	if cmd == nil {
		err = utils.ErrorCommand
		return
	}

	cmd.ExecutedDate = time.Now()

	// validate argument
	if cmd.Arguments.Amount > 1000 {
		err = utils.ErrorMaximumAmountTransfer
		return
	}

	if cmd.Arguments.Amount < 1 {
		err = utils.ErrorMinimumAmountTransfer
		return
	}

	if cmd.Arguments.ReferenceNumber == "" {
		err = utils.ErrorReferenceNumber
		return
	}

	// get from user
	userFrom, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.From)

	if err != nil {
		return
	}

	if userFrom.Balance < cmd.Arguments.Amount {
		err = utils.SetErrorInsufficient(userFrom.Currency, userFrom.Balance)
		return
	}

	// get from to
	userTo, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.To)

	if err != nil {
		return
	}

	// modify balance from user
	balance := userFrom.Balance - cmd.Arguments.Amount
	err = pl.repo.UpdateUserBalance(userFrom.Id, balance)
	if err != nil {
		return
	}

	// modify balance to user
	balance = userTo.Balance + cmd.Arguments.Amount
	err = pl.repo.UpdateUserBalance(userTo.Id, balance)
	if err != nil {
		return
	}

	//re login to update logged user
	err = pl.repo.Login(userFrom.Id)

	return
}
