package schemas

import "time"

type AtmSimulation interface {
	Execute(cmd *Command) error
}

type AtmSimulationView interface {
	StartDisplay(cmd *Command) (err error)
	EndDisplay(cmd *Command) (err error)
}

type Command struct {
	Command      int
	Arguments    Arguments
	Service      AtmSimulation
	View         AtmSimulationView
	ExecutedDate time.Time
}

type Arguments struct {
	AtmMachineArg AtmMachineArguments
	CsvArg        CsvArguments
}

type AtmMachineArguments struct {
	To              string
	From            string
	Amount          int64
	Pin             string
	ReferenceNumber string
}

type CsvArguments struct {
	Chosen int
	Path   string
}
