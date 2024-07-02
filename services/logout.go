package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type logout struct {
	repo datasource.Datasources
}

func NewLogout(d datasource.Datasources) *logout {
	pl := &logout{d}
	return pl
}

func (pl *logout) Execute(cmd *schemas.Command) (err error) {

	err = pl.repo.Logout()

	return
}
