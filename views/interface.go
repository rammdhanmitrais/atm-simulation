package views

import (
	"atm-simulation/schemas"
)

type Views interface {
	Transaction() schemas.AtmSimulationView
	Withdraw() schemas.AtmSimulationView
	ViewBalance() schemas.AtmSimulationView
	FundTransfer() schemas.AtmSimulationView
	Login() schemas.AtmSimulationView
	Logout() schemas.AtmSimulationView
	ReadCsv() schemas.AtmSimulationView
}
