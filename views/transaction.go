package views

import (
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type transaction struct {
	*views
}

func NewTransaction(s *views) *transaction {
	pl := &transaction{s}
	return pl
}

func (pl *transaction) StartDisplay(cmd *schemas.Command) (err error) {

	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Println()
	fmt.Printf("1 %s \n", utils.Withdraw)
	fmt.Printf("2 %s \n", utils.ViewBalance)
	fmt.Printf("3 %s \n", utils.FundTransfer)
	fmt.Printf("4 %s \n", utils.Exit)
	fmt.Print("Please choose option[4]: ")

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	if ascii == 27 || ascii == 3 {
		err = utils.ErrorExit
	}

	switch input {
	case 1:
		cmd.Command = utils.WithdrawCommand
	case 2:
		cmd.Command = utils.ViewBalanceCommand
	case 3:
		cmd.Command = utils.FundTransferCommand
	case 4:
		cmd.Command = utils.LogoutCommand
	default:
		cmd.Command = utils.LogoutCommand
	}

	return
}

func (pl *transaction) EndDisplay(cmd *schemas.Command) (err error) {
	return
}
