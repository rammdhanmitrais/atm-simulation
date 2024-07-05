package datasource

import (
	"atm-simulation/utils"
)

type datasource struct{}

func NewDatasource() *datasource {
	ds := new(datasource)

	return ds
}

func (d datasource) GetUserByAccountNumber(accountNumber string) (result User, err error) {
	for _, user := range userAccounts {
		if user.AccountNumber == accountNumber {
			result = user
			return
		}
	}

	err = utils.ErrorInvalidAccount

	return
}

func (d datasource) GetLoggedUser() (user User, err error) {
	if LoggedUser == nil {
		err = utils.ErrorUserLogged
		return
	}

	user = *LoggedUser
	return
}

func (d datasource) UpdateUserBalance(id int, balance int64) (err error) {
	if id > len(userAccounts) || id < 0 {
		err = utils.ErrorInvalidAccount
		return
	}

	for i, user := range userAccounts {
		if user.Id == id {
			userAccounts[i].Balance = balance
			return
		}
	}

	err = utils.ErrorInvalidAccount
	return
}

func (d datasource) Login(id int) (err error) {
	for _, user := range userAccounts {
		if user.Id == id {
			LoggedUser = &user
			return
		}
	}

	err = utils.ErrorInvalidAccount

	return
}

func (d datasource) Logout() (err error) {
	LoggedUser = nil

	return
}
