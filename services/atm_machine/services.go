package atm_machine

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
)

type ServiceDatasources struct {
	UserDatasource        datasource.UserDatasources
	TransactionDatasource datasource.TransactionDatasources
}

type atmMachineServices struct {
	withdraw           schemas.AtmSimulation
	viewBalance        schemas.AtmSimulation
	fundTransfer       schemas.AtmSimulation
	login              schemas.AtmSimulation
	logout             schemas.AtmSimulation
	transactionHistory schemas.AtmSimulation
}

func NewAtmMachineService(d ServiceDatasources) *atmMachineServices {
	st := new(atmMachineServices)

	st.withdraw = NewWithdraw(d)
	st.login = NewLogin(d)
	st.logout = NewLogout(d)
	st.viewBalance = NewViewBalance(d)
	st.fundTransfer = NewFundTransfer(d)
	st.transactionHistory = NewTransactionHistory(d)

	return st
}

func (s atmMachineServices) Withdraw() schemas.AtmSimulation {
	return s.withdraw
}

func (s atmMachineServices) ViewBalance() schemas.AtmSimulation {
	return s.viewBalance
}

func (s atmMachineServices) FundTransfer() schemas.AtmSimulation {
	return s.fundTransfer
}

func (s atmMachineServices) Login() schemas.AtmSimulation {
	return s.login
}

func (s atmMachineServices) Logout() schemas.AtmSimulation {
	return s.logout
}

func (s atmMachineServices) TransactionHistory() schemas.AtmSimulation {
	return s.transactionHistory
}
