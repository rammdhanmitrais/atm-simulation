package views

import (
	"atm-simulation/schemas"
)

type views struct {
	withdraw     schemas.AtmSimulationView
	viewBalance  schemas.AtmSimulationView
	fundTransfer schemas.AtmSimulationView
	login        schemas.AtmSimulationView
	logout       schemas.AtmSimulationView
	transaction  schemas.AtmSimulationView
	readCsv      schemas.AtmSimulationView
}

func NewView() *views {
	st := new(views)

	st.withdraw = NewWithdraw(st)
	st.login = NewLogin(st)
	st.logout = NewLogout(st)
	st.viewBalance = NewViewBalance(st)
	st.transaction = NewTransaction(st)
	st.fundTransfer = NewFundTransfer(st)
	st.readCsv = NewReadCsv(st)

	return st
}

func (s views) Withdraw() schemas.AtmSimulationView {
	return s.withdraw
}

func (s views) ViewBalance() schemas.AtmSimulationView {
	return s.viewBalance
}

func (s views) Login() schemas.AtmSimulationView {
	return s.login
}

func (s views) Logout() schemas.AtmSimulationView {
	return s.logout
}

func (s views) Transaction() schemas.AtmSimulationView {
	return s.transaction
}

func (s views) FundTransfer() schemas.AtmSimulationView {
	return s.fundTransfer
}

func (s views) ReadCsv() schemas.AtmSimulationView {
	return s.readCsv
}
