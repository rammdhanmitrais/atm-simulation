package atm_machine

import (
	"atm-simulation/schemas"
)

type logout struct {
	repo ServiceDatasources
}

func NewLogout(d ServiceDatasources) *logout {
	pl := &logout{d}
	return pl
}

func (pl *logout) Execute(cmd *schemas.Command) (err error) {

	err = pl.repo.UserDatasource.Logout()

	return
}
