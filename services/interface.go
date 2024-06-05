package services

import (
	"atm-simulation/schemas"
)

type Services interface {
	Withdraw() schemas.AtmMachine
	ViewBalance() schemas.AtmMachine
	FundTransfer() schemas.AtmMachine
	Login() schemas.AtmMachine
	Logout() schemas.AtmMachine
}
