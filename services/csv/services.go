package csv

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type csvServices struct {
	readCsv schemas.AtmSimulation
}

func NewCsvService(d datasource.UserDatasources) *csvServices {
	st := new(csvServices)

	st.readCsv = NewReadCsv(d)

	return st
}

func (s csvServices) ReadCsv() schemas.AtmSimulation {
	return s.readCsv
}
