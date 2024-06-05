package services

import (
	"atm-simulation/schemas"
)

type viewBalance struct {
	*services
}

func NewViewBalance(s *services) *viewBalance {
	pl := &viewBalance{s}
	return pl
}

func (pl *viewBalance) Execute(cmd *schemas.Command) (err error) {

	return
}