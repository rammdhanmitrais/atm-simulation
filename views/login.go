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
		err = fmt.Errorf("account number should only contains numbers") 
		return
	}

	if len(account) != 6 {
		err = fmt.Errorf("account number should have 6 digits length") 
		return
	}

	fmt.Print("Enter Pin: ")
	pin, _ := reader.ReadString('\n')
	pin = strings.Trim(string(pin), "\n")

	if !utils.ValidateIsContainNumberOnly(pin) {
		err = fmt.Errorf("PIN should only contains numbers") 
		return
	}

	if len(pin) != 6 {
		err = fmt.Errorf("PIN should have 6 digits length") 
		return
	}

	cmd.Arguments.From = account
	cmd.Arguments.Pin = pin

	return
}

func (pl *login) EndDisplay(cmd *schemas.Command) (err error){
	fmt.Println()
	fmt.Println("Login successfully")
	cmd.Command = utils.TransactionCommand

	return
}
