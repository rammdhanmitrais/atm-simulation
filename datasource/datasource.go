package datasource

type datasource struct {
	UserDatasource        *userDatasource
	TransactionDatasource *transactionDatasource
}

func NewDatasource() *datasource {
	ds := new(datasource)
	ds.UserDatasource = NewUserDatasource()
	ds.TransactionDatasource = NewTransactionDatasource()
	return ds
}
