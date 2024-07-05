package datasource

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
