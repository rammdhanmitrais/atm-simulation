package utils

const LayoutDateTime = "2006-01-02 15:04:05"

const (
	Withdraw           = "Withdraw"
	FundTransfer       = "Fund Transfer"
	ViewBalance        = "View Balance"
	TransactionHistory = "Transactions History"
	Exit               = "Exit"
	Back               = "Back"
)

const (
	Credit = "Credit"
	Debit  = "Debit"
)

const (
	LoginCommand              = 0
	TransactionCommand        = 1
	WithdrawCommand           = 2
	ViewBalanceCommand        = 3
	FundTransferCommand       = 4
	ChangeDatasourceCommand   = 5
	TransactionHistoryCommand = 6
	LogoutCommand             = 99
)

var WithdrawValues = [3]int64{10, 50, 100}
