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
	cmd.Arguments.From = strings.Trim(account, "\n")
	if len(cmd.Arguments.From) != 6 {
		err = fmt.Errorf("account number should have 6 digits length") 
		return
	}

	fmt.Print("Enter Pin: ")
	pin, _ := reader.ReadString('\n')
	cmd.Arguments.Pin = strings.Trim(string(pin), "\n")

	if len(cmd.Arguments.Pin) != 6 {
		err = fmt.Errorf("PIN should have 6 digits length") 
		return
	}

	return
}

func (pl *login) EndDisplay(cmd *schemas.Command) (err error){
	fmt.Println()
	fmt.Println("Login successfully")
	cmd.Command = utils.TransactionCommand

	return
}
