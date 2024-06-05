package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type logout struct {
	*services
}

func NewLogout(s *services) *logout {
	pl := &logout{s}
	return pl
}

func (pl *logout) Execute(cmd *schemas.Command) (err error) {
	
	datasource.LoggedUser = nil

	return
}