package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type viewBalance struct {
	repo datasource.Datasources
}

func NewViewBalance(d datasource.Datasources) *viewBalance {
	pl := &viewBalance{d}
	return pl
}

func (pl *viewBalance) Execute(cmd *schemas.Command) (err error) {

	return
}
