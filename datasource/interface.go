package datasource

type UserDatasources interface {
	GetUserByAccountNumber(accountNumber string) (result User, err error)
	GetLoggedUser() (user User, err error)
	UpdateUserBalance(id int, balance int64) (err error)
	Login(id int) (err error)
	Logout() (err error)
}
