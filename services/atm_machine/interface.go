package atm_machine

import (
	"atm-simulation/schemas"
)

type AtmMachineServices interface {
	Withdraw() schemas.AtmSimulation
	ViewBalance() schemas.AtmSimulation
	FundTransfer() schemas.AtmSimulation
	Login() schemas.AtmSimulation
	Logout() schemas.AtmSimulation
}
