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

type fundTransfer struct {
	*views
}

func NewFundTransfer(s *views) *fundTransfer {
	pl := &fundTransfer{s}
	return pl
}

func (pl *fundTransfer) StartDisplay(cmd *schemas.Command) (err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Please enter [destination account] or ")
	fmt.Print("Press 0 to go back to Transaction: ")

	destination, _ := reader.ReadString('\n')
	destination = strings.Trim(destination, "\n")

	if destination == "0" {
		cmd.Command = utils.TransactionCommand
		return
	}

	if !utils.ValidateIsContainNumberOnly(destination) {
		err = utils.ErrorInvalidAccount
		return
	}

	amount, command, err := fundTransferGetAmount()
	if command > 0 {
		cmd.Command = utils.TransactionCommand
		return
	}

	if err != nil {
		return
	}

	reference, command := fundTransferGetReference()
	if command > 0 {
		cmd.Command = utils.TransactionCommand
		return
	}

	fmt.Println()
	fmt.Println("Transfer Confirmation")
	fmt.Printf("Destination Account: %s\n", destination)
	fmt.Printf("Transfer Amount: $%d\n", amount)
	fmt.Printf("Reference Number: %s\n", reference)

	fmt.Println("1 Confirm Transfer")
	fmt.Println("2 Cancel Transfer")

	fmt.Print("Please choose option[2]: ")

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	if ascii == 27 || ascii == 3 {
		err = utils.ErrorExit
		return
	}

	if input != 1 {
		cmd.Command = utils.LogoutCommand
	}

	cmd.Arguments.AtmMachineArg.From = datasource.LoggedUser.AccountNumber
	cmd.Arguments.AtmMachineArg.To = destination
	cmd.Arguments.AtmMachineArg.Amount = int64(amount)
	cmd.Arguments.AtmMachineArg.ReferenceNumber = reference

	return
}

func fundTransferGetAmount() (response int, command int, err error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Please enter [transfer amount] or")
	fmt.Print("Press 0 to go back to Transaction: ")

	input, _ := reader.ReadString('\n')
	if input == "0" {
		command = utils.TransactionCommand
	}

	input = strings.Trim(input, "\n")
	response, err = strconv.Atoi(input)

	if err != nil {
		err = utils.ErrorInvalidAccount
	}

	return
}

func fundTransferGetReference() (response string, command int) {

	reader := bufio.NewReader(os.Stdin)
	reference := utils.GenerateNDigitRandom(6)

	fmt.Println()
	fmt.Printf("Reference Number: %s\n", reference)
	fmt.Println("Please enter to continue or")
	fmt.Print("Press 0 to go back to Transaction: ")

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	if input == 0 && ascii != 10 {
		command = utils.TransactionCommand
	}

	response = reference

	return
}

func (pl *fundTransfer) EndDisplay(cmd *schemas.Command) (err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Fund Transfer Summary")
	fmt.Printf("Destination Account: %s\n", cmd.Arguments.AtmMachineArg.To)
	fmt.Printf("Transfer Amount: $%d\n", cmd.Arguments.AtmMachineArg.Amount)
	fmt.Printf("Reference Number: %s\n", cmd.Arguments.AtmMachineArg.ReferenceNumber)

	fmt.Println("1 Transaction")
	fmt.Println("2 Exit")

	fmt.Print("Please choose option[2]: ")

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	if ascii == 27 || ascii == 3 {
		err = utils.ErrorExit
		return
	}

	if input != 1 {
		cmd.Command = utils.LogoutCommand
	} else {
		cmd.Command = utils.TransactionCommand
	}

	return
}
