package csv

import "atm-simulation/schemas"

type CsvServices interface {
	ReadCsv() schemas.AtmSimulation
}
