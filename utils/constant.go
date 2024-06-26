package utils

const LayoutDateTime = "2006-01-02 15:04:05"

const (
	Withdraw 		= "Withdraw"
	FundTransfer 	= "Fund Transfer"
	ViewBalance 	= "View Balance"
	Exit 			= "Exit"
	Back 			= "Back"
)

const (
	LoginCommand 		= 0
	TransactionCommand 	= 1
	WithdrawCommand 	= 2
	ViewBalanceCommand 	= 3
	FundTransferCommand = 4
	LogoutCommand 		= 99
)

var WithdrawValues = []int64{10, 50, 100}