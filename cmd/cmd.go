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
	userDatasource := datasource.NewUserDatasource()
	transactionDatasource := datasource.NewTransactionDatasource()
	serviceDatasource := atm_machine.ServiceDatasources{
		UserDatasource:        userDatasource,
		TransactionDatasource: transactionDatasource,
	}

	atmMachineService = atm_machine.NewAtmMachineService(serviceDatasource)
	csvServices = csv.NewCsvService(userDatasource)
	View = views.NewView()

	var command *int
	for {
		result, err := Process(command)
		if err != nil {
			fmt.Println(err.Error())

			if result != utils.ChangeDatasourceCommand {
				result = utils.LoginCommand
			}
		}

		command = &result
	}
}
