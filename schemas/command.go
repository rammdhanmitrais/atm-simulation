package schemas

import "time"

type AtmMachine interface {
	Execute(cmd *Command) error
}

type AtmMachineView interface {
	StartDisplay(cmd *Command) (err error)
	EndDisplay(cmd *Command) (err error)
}

type Command struct {
	Command      int
	Arguments    Arguments
	Service      AtmMachine
	View         AtmMachineView
	ExecutedDate time.Time
}

type Arguments struct {
	To              string
	From            string
	Amount          int64
	Pin             string
	ReferenceNumber string
}
