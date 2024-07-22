package atm_machine

import (
	"atm-simulation/schemas"
)

type viewBalance struct {
	repo ServiceDatasources
}

func NewViewBalance(d ServiceDatasources) *viewBalance {
	pl := &viewBalance{d}
	return pl
}

func (pl *viewBalance) Execute(cmd *schemas.Command) (err error) {

	return
}
