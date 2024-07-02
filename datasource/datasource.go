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
		Name:          "John Doe",
		Pin:           "012108",
		Currency:      "$",
		Balance:       100,
		AccountNumber: "112233",
	},
	{
		Name:          "Jane Doe",
		Pin:           "932012",
		Currency:      "$",
		Balance:       30,
		AccountNumber: "112244",
	},
}

type Datasources interface {
	GetUserByAccountNumber(accountNumber string) (result User, err error)
	GetLoggedUser() (user User, err error)
	UpdateUserBalance(id int, balance int64) (err error)
	Login(id int) (err error)
	Logout() (err error)
}

type datasource struct {
	userRepository Datasources
}

func NewDatasource() *datasource {
	ds := new(datasource)

	ds.userRepository = NewUserRepository(ds)

	return ds
}

func (d datasource) GetUserByAccountNumber(accountNumber string) (result User, err error) {
	return d.userRepository.GetUserByAccountNumber(accountNumber)
}

func (d datasource) GetLoggedUser() (user User, err error) {
	return d.userRepository.GetLoggedUser()
}

func (d datasource) UpdateUserBalance(id int, balance int64) (err error) {
	return d.userRepository.UpdateUserBalance(id, balance)
}

func (d datasource) Login(id int) (err error) {
	return d.userRepository.Login(id)
}

func (d datasource) Logout() (err error) {
	return d.userRepository.Logout()
}
