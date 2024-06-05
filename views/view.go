package views

import (
	"atm-simulation/schemas"
)

type views struct {
	withdraw schemas.AtmMachineView
	viewBalance schemas.AtmMachineView
	fundTransfer schemas.AtmMachineView
	login schemas.AtmMachineView
	logout schemas.AtmMachineView
	transaction schemas.AtmMachineView
}

func NewView() *views {
	st := new(views)

	st.withdraw = NewWithdraw(st)
	st.login = NewLogin(st)
	st.logout = NewLogout(st)
	st.viewBalance = NewViewBalance(st)
	st.transaction = NewTransaction(st)
	st.fundTransfer = NewFundTransfer(st)

	return st
}

func (s views) Withdraw() schemas.AtmMachineView {
	return s.withdraw
}

func (s views) ViewBalance() schemas.AtmMachineView {
	return s.viewBalance
}

func (s views) Login() schemas.AtmMachineView {
	return s.login
}

func (s views) Logout() schemas.AtmMachineView {
	return s.logout
}

func (s views) Transaction() schemas.AtmMachineView {
	return s.transaction
}

func (s views) FundTransfer() schemas.AtmMachineView {
	return s.fundTransfer
}