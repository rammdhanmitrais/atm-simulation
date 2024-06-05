package datasource

import "errors"

var UserAccounts []User
var LoggedUser *User

type User struct {
	Index			int
	Name 			string
	Pin 			string
	Balance 		int64
	Currency 		string
	AccountNumber 	string
}

func NewUser(){
	UserAccounts = []User{
		{
			Name: "John Doe",
			Pin: "012108",
			Currency: "$",
			Balance: 100,
			AccountNumber: "112233",
		},
		{
			Name: "Jane Doe",
			Pin: "932012",
			Currency: "$",
			Balance: 30,
			AccountNumber: "112244",
		},
	}
}

func GetUserByAccountNumber(accountNumber string) (result User, err error) {
	
	for i, user := range UserAccounts {
		if user.AccountNumber == accountNumber {
			user.Index = i
			result = user
			return
		}
	}

	err = errors.New("invalid account")

	return
}