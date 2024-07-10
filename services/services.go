package services

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type services struct {
	withdraw     schemas.AtmMachine
	viewBalance  schemas.AtmMachine
	fundTransfer schemas.AtmMachine
	login        schemas.AtmMachine
	logout       schemas.AtmMachine
}

func NewService(d datasource.UserDatasources) *services {
	st := new(services)

	st.withdraw = NewWithdraw(d)
	st.login = NewLogin(d)
	st.logout = NewLogout(d)
	st.viewBalance = NewViewBalance(d)
	st.fundTransfer = NewFundTransfer(d)

	return st
}

func (s services) Withdraw() schemas.AtmMachine {
	return s.withdraw
}

func (s services) ViewBalance() schemas.AtmMachine {
	return s.viewBalance
}

func (s services) FundTransfer() schemas.AtmMachine {
	return s.fundTransfer
}

func (s services) Login() schemas.AtmMachine {
	return s.login
}

func (s services) Logout() schemas.AtmMachine {
	return s.logout
}
