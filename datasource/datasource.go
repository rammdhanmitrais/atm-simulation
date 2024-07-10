package datasource

type datasource struct {
	UserDataSource *userDatasource
}

func NewDatasource() *datasource {
	ds := new(datasource)
	ds.UserDataSource = NewUserDatasource()
	return ds
}
