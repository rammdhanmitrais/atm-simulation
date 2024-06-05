package views

import (
	"atm-simulation/schemas"
)

type Views interface {
	Transaction() schemas.AtmMachineView
	Withdraw() schemas.AtmMachineView
	ViewBalance() schemas.AtmMachineView
	FundTransfer() schemas.AtmMachineView
	Login() schemas.AtmMachineView
	Logout() schemas.AtmMachineView
}
