package cmd

import (
	"atm-simulation/datasource"
	"atm-simulation/services"
	"atm-simulation/utils"
	"atm-simulation/views"
	"fmt"
)

var Service services.Services
var View views.Views

func Start() {
	datasource := datasource.NewDatasource()
	Service = services.NewService(datasource)
	View = views.NewView()

	var command *int
	for {

		result, err := Process(command)
		if err != nil {
			fmt.Println(err.Error())
			result = utils.LoginCommand
		}

		command = &result
	}
}
