package views

import (
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type login struct {
	*views
}

func NewLogin(s *views) *login {
	pl := &login{s}
	return pl
}

func (pl *login) StartDisplay(cmd *schemas.Command) (err error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Print("Enter Account Number: ")
	account, _ := reader.ReadString('\n')
	account = strings.Trim(account, "\n")

	if !utils.ValidateIsContainNumberOnly(account) {
		err = utils.ErrorAccountNumberNotNumber
		return
	}

	if len(account) != 6 {
		err = utils.ErrorAccountNumberDigit
		return
	}

	fmt.Print("Enter Pin: ")
	pin, _ := reader.ReadString('\n')
	pin = strings.Trim(string(pin), "\n")

	if !utils.ValidateIsContainNumberOnly(pin) {
		err = utils.ErrorPINNotNumber
		return
	}

	if len(pin) != 6 {
		err = utils.ErrorPINDigit
		return
	}

	cmd.Arguments.AtmMachineArg.From = account
	cmd.Arguments.AtmMachineArg.Pin = pin

	return
}

func (pl *login) EndDisplay(cmd *schemas.Command) (err error) {
	fmt.Println()
	fmt.Println("Login successfully")
	cmd.Command = utils.TransactionCommand

	return
}
