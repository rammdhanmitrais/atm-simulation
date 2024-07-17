package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type logout struct {
	repo datasource.UserDatasources
}

func NewLogout(d datasource.UserDatasources) *logout {
	pl := &logout{d}
	return pl
}

func (pl *logout) Execute(cmd *schemas.Command) (err error) {

	err = pl.repo.Logout()

	return
}
