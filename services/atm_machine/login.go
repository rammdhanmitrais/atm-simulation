package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"time"
)

type login struct {
	repo datasource.UserDatasources
}

func NewLogin(d datasource.UserDatasources) *login {
	pl := &login{d}
	return pl
}

func (pl *login) Execute(cmd *schemas.Command) (err error) {
	if cmd == nil {
		err = utils.ErrorCommand
		return
	}

	cmd.ExecutedDate = time.Now()

	// get user
	user, err := pl.repo.GetUserByAccountNumber(cmd.Arguments.AtmMachineArg.From)
	if err != nil {
		return
	}

	if user.AccountNumber == "" || user.Pin != cmd.Arguments.AtmMachineArg.Pin {
		err = utils.ErrorInvalidAccountPin
		return
	}

	//login to update logged user
	err = pl.repo.Login(user.Id)

	return
}
