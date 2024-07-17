package cmd

import (
	"atm-simulation/datasource"
	"atm-simulation/services/atm_machine"
	"atm-simulation/services/csv"
	"atm-simulation/utils"
	"atm-simulation/views"
	"fmt"
)

var atmMachineService atm_machine.AtmMachineServices
var csvServices csv.CsvServices
var View views.Views

func Start() {
	datasource := datasource.NewUserDatasource()
	atmMachineService = atm_machine.NewAtmMachineService(datasource)
	csvServices = csv.NewCsvService(datasource)
	View = views.NewView()

	var command *int
	for {
		result, err := Process(command)
		if err != nil {
			fmt.Println(err.Error())
			result = utils.FirstCommand
		}

		command = &result
	}
}
