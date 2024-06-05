package views

import (
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"fmt"
)

type logout struct {
	*views
}

func NewLogout(s *views) *logout {
	pl := &logout{s}
	return pl
}

func (pl *logout) StartDisplay(cmd *schemas.Command) (err error) {
	return
}

func (pl *logout) EndDisplay(cmd *schemas.Command)(err error){
	fmt.Println()
	fmt.Println("logout successfully")
	cmd.Command = utils.LoginCommand

	return
}
