package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type viewBalance struct {
	repo datasource.UserDatasources
}

func NewViewBalance(d datasource.UserDatasources) *viewBalance {
	pl := &viewBalance{d}
	return pl
}

func (pl *viewBalance) Execute(cmd *schemas.Command) (err error) {

	return
}
