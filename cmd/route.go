package cmd

import (
	"atm-simulation/schemas"
	"atm-simulation/utils"
)

func Process(input *int) (command int, err error) {
	var cmd = new(schemas.Command)

	if input == nil {
		cmd.Command = utils.ChangeDatasourceCommand
		input = &cmd.Command
	} else {
		cmd.Command = *input
	}

	// set service and view based on command
	setAtmMachineConnection(cmd)

	// start display view
	err = cmd.View.StartDisplay(cmd)
	if err != nil {
		return
	}

	// if command with input diffence it means there is change view in start display
	if cmd.Command != *input {
		command = cmd.Command
		return
	}

	if cmd.Service != nil {
		// execute the service
		err = cmd.Service.Execute(cmd)
		if err != nil {
			command = cmd.Command
			return
		}
	}

	// end display view
	err = cmd.View.EndDisplay(cmd)
	command = cmd.Command

	return
}

func setAtmMachineConnection(command *schemas.Command) {
	switch command.Command {
	case utils.LoginCommand:
		command.Service = atmMachineService.Login()
		command.View = View.Login()
	case utils.TransactionCommand:
		command.View = View.Transaction()
	case utils.WithdrawCommand:
		command.Service = atmMachineService.Withdraw()
		command.View = View.Withdraw()
	case utils.ViewBalanceCommand:
		command.Service = atmMachineService.ViewBalance()
		command.View = View.ViewBalance()
	case utils.FundTransferCommand:
		command.Service = atmMachineService.FundTransfer()
		command.View = View.FundTransfer()
	case utils.LogoutCommand:
		command.Service = atmMachineService.Logout()
		command.View = View.Logout()
	case utils.ChangeDatasourceCommand:
		command.Service = csvServices.ReadCsv()
		command.View = View.ReadCsv()
	case utils.TransactionHistoryCommand:
		command.Service = atmMachineService.TransactionHistory()
		command.View = View.TransactionHistory()
	}
}
