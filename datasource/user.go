package datasource

import "atm-simulation/utils"

type userRepository struct {
	*datasource
}

func NewUserRepository(d *datasource) *userRepository {
	ds := &userRepository{d}
	return ds
}

func (d *userRepository) GetUserByAccountNumber(accountNumber string) (result User, err error) {

	for i, user := range userAccounts {
		if user.AccountNumber == accountNumber {
			user.Id = i
			result = user
			return
		}
	}

	err = utils.ErrorInvalidAccount

	return
}

func (d *userRepository) UpdateUserBalance(id int, balance int64) (err error) {
	if id > len(userAccounts) {
		err = utils.ErrorInvalidAccount
		return
	}

	userAccounts[id].Balance = balance

	return
}

func (d *userRepository) GetLoggedUser() (user User, err error) {

	if LoggedUser == nil {
		err = utils.ErrorUserLogged
		return
	}

	user = *LoggedUser
	return
}

func (d *userRepository) Login(id int) (err error) {

	for _, user := range userAccounts {
		if user.Id == id {
			LoggedUser = &user
			return
		}
	}

	err = utils.ErrorInvalidAccount

	return
}

func (d *userRepository) Logout() (err error) {

	LoggedUser = nil

	return
}
