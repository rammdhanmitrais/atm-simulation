package services

import (
	"atm-simulation/schemas"
)

type services struct {
	withdraw schemas.AtmMachine
	viewBalance schemas.AtmMachine
	fundTransfer schemas.AtmMachine
	login schemas.AtmMachine
	logout schemas.AtmMachine
}

func NewService() *services {
	st := new(services)

	st.withdraw = NewWithdraw(st)
	st.login = NewLogin(st)
	st.logout = NewLogout(st)
	st.viewBalance = NewViewBalance(st)
	st.fundTransfer = NewFundTransfer(st)

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