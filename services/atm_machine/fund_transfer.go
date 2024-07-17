package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"time"
)

type fundTransfer struct {
	repo datasource.UserDatasources
}

func NewFundTransfer(d datasource.UserDatasources) *fundTransfer {
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
	userFrom, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.AtmMachineArg.From)

	if err != nil {
		return
	}

	if userFrom.Balance < cmd.Arguments.AtmMachineArg.Amount {
		err = utils.SetErrorInsufficient(userFrom.Currency, userFrom.Balance)
		return
	}

	// get from to
	userTo, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.AtmMachineArg.To)

	if err != nil {
		return
	}

	// modify balance from user
	balance := userFrom.Balance - cmd.Arguments.AtmMachineArg.Amount
	err = pl.repo.UpdateUserBalance(userFrom.Id, balance)
	if err != nil {
		return
	}

	// modify balance to user
	balance = userTo.Balance + cmd.Arguments.AtmMachineArg.Amount
	err = pl.repo.UpdateUserBalance(userTo.Id, balance)
	if err != nil {
		return
	}

	//re login to update logged user
	err = pl.repo.Login(userFrom.Id)

	return
}
