package views

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type withdraw struct {
	*views
}

func NewWithdraw(s *views) *withdraw {
	pl := &withdraw{s}
	return pl
}

func (pl *withdraw) StartDisplay(cmd *schemas.Command) (err error) {

	reader := bufio.NewReaderSize(os.Stdin, 1)

	idx := 1
	fmt.Println()
	for _, v := range utils.WithdrawValues {
		fmt.Printf("%d $%d \n", idx, v)
		idx++
	}

	otherIdx := idx
	fmt.Printf("%d Other\n", otherIdx)

	backIdx := otherIdx + 1
	fmt.Printf("%d Back\n", backIdx)

	fmt.Printf("Please choose option[%d]: ", backIdx)

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	var value int64
	if input > len(utils.WithdrawValues)+2 || input < 1 || input == backIdx {
		cmd.Command = utils.TransactionCommand
		return
	} else if input == otherIdx {
		val, errors := otherScreen()
		if errors != nil {
			err = errors
			return
		}

		value = int64(val)
	} else {
		value = utils.WithdrawValues[input-1]
	}

	cmd.Arguments.From = datasource.LoggedUser.AccountNumber
	cmd.Arguments.Amount = value

	return
}

func otherScreen() (response int, err error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Other Withdraw")
	fmt.Print("Enter amount to withdraw: ")

	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")

	response, err = strconv.Atoi(input)

	if err != nil {
		err = utils.ErrorInvalidAmount
		return
	}

	// validate amount should multiple of 10
	if response%10 != 0 {
		err = utils.ErrorInvalidAmount
		return
	}

	// validate amount should less than 1000
	if response > 1000 {
		err = utils.ErrorMaximumAmountTransfer
		return
	}

	return
}

func (pl *withdraw) EndDisplay(cmd *schemas.Command) (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Println()
	fmt.Println("Summary")
	fmt.Printf("Date: %s\n", cmd.ExecutedDate.Format(utils.LayoutDateTime))
	fmt.Printf("Withdraw: $%d\n", cmd.Arguments.Amount)
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
