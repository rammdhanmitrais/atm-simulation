package views

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type transactionHistory struct {
	*views
}

func NewTransactionHistory(s *views) *transactionHistory {
	pl := &transactionHistory{s}
	return pl
}

func (pl *transactionHistory) StartDisplay(cmd *schemas.Command) (err error) {

	return
}

func (pl *transactionHistory) EndDisplay(cmd *schemas.Command) (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Transaction Type", "Transaction Date", "Account Number", "From / To", "Initial Balance", "Amount", "Status"})

	for _, data := range datasource.LoggedUser.TransactionHistory {
		if data.Type != utils.FundTransfer {
			data.FromToAccountNumber = "-"
		}

		row := []string{data.Type, data.TransactionDate.Format(utils.LayoutDateTime), data.AccountNumber, data.FromToAccountNumber, fmt.Sprintf("$%d", data.InitialBalance), fmt.Sprintf("$%d", data.Amount), data.CreditOrDebit}
		table.Append(row)
	}

	table.Render()

	fmt.Println("1 Transaction")
	fmt.Println("2 Exit")

	fmt.Print("Please choose option[2]: ")

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	if ascii == 27 || ascii == 3 {
		err = utils.ErrorExit
		return
	}

	if input == 1 {
		cmd.Command = utils.TransactionCommand
	} else {
		cmd.Command = utils.LogoutCommand
	}

	return
}
