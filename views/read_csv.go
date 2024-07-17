package views

import (
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type readCsv struct {
	*views
}

func NewReadCsv(s *views) *readCsv {
	pl := &readCsv{s}
	return pl
}

func (pl *readCsv) StartDisplay(cmd *schemas.Command) (err error) {
	reader := bufio.NewReaderSize(os.Stdin, 1)

	fmt.Println()
	fmt.Println("Choose Datasource")
	fmt.Println("1 Default Datasource")
	fmt.Println("2 Read CSV")

	fmt.Print("Please choose option[1]: ")

	ascii, _ := reader.ReadByte()
	input, _ := strconv.Atoi(string(ascii))

	if ascii == 27 || ascii == 3 {
		err = utils.ErrorExit
		return
	}

	cmd.Arguments.CsvArg.Chosen = input
	if cmd.Arguments.CsvArg.Chosen == 1 {
		return
	}

	path := inputCsvPath()
	cmd.Arguments.CsvArg.Path = path

	return
}

func inputCsvPath() (path string) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("Please enter csv path: ")

	path, _ = reader.ReadString('\n')
	path = strings.Trim(path, "\n")

	return
}

func (pl *readCsv) EndDisplay(cmd *schemas.Command) (err error) {
	cmd.Command = utils.LoginCommand
	return
}
