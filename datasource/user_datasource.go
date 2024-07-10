package datasource

import "atm-simulation/utils"

type User struct {
	Id            int
	Name          string
	Pin           string
	Balance       int64
	Currency      string
	AccountNumber string
}

var LoggedUser *User
var userAccounts = [2]User{
	{
		Id:            1,
		Name:          "John Doe",
		Pin:           "012108",
		Currency:      "$",
		Balance:       100,
		AccountNumber: "112233",
	},
	{
		Id:            2,
		Name:          "Jane Doe",
		Pin:           "932012",
		Currency:      "$",
		Balance:       30,
		AccountNumber: "112244",
	},
}

type userDatasource struct{}

func NewUserDatasource() *userDatasource {
	return &userDatasource{}
}

func (d userDatasource) GetUserByAccountNumber(accountNumber string) (result User, err error) {
	for _, user := range userAccounts {
		if user.AccountNumber == accountNumber {
			result = user
			return
		}
	}

	err = utils.ErrorInvalidAccount

	return
}

func (d userDatasource) GetLoggedUser() (user User, err error) {
	if LoggedUser == nil {
		err = utils.ErrorUserLogged
		return
	}

	user = *LoggedUser
	return
}

func (d userDatasource) UpdateUserBalance(id int, balance int64) (err error) {
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

func (d userDatasource) Login(id int) (err error) {
	for _, user := range userAccounts {
		if user.Id == id {
			LoggedUser = &user
			return
		}
	}

	err = utils.ErrorInvalidAccount

	return
}

func (d userDatasource) Logout() (err error) {
	LoggedUser = nil

	return
}
