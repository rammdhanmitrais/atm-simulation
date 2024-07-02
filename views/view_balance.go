package views

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type viewBalance struct {
	*views
}

func NewViewBalance(s *views) *viewBalance {
	pl := &viewBalance{s}
	return pl
}

func (pl *viewBalance) StartDisplay(cmd *schemas.Command) (err error) {

	return
}

func (pl *viewBalance) EndDisplay(cmd *schemas.Command) (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Println()
	fmt.Println("Account Info")
	fmt.Printf("Date: %s\n", time.Now().Format(utils.LayoutDateTime))
	fmt.Printf("Account Number: %s\n", datasource.LoggedUser.AccountNumber)
	fmt.Printf("Balance: $%d\n", datasource.LoggedUser.Balance)

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
