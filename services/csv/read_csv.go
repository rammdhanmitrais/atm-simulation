package csv

import (
	"atm-simulation/datasource"
	"atm-simulation/schemas"
	"atm-simulation/utils"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type readCsv struct {
	repo datasource.UserDatasources
}

func NewReadCsv(d datasource.UserDatasources) *readCsv {
	pl := &readCsv{d}
	return pl
}

func (pl *readCsv) Execute(cmd *schemas.Command) (err error) {

	if cmd.Arguments.CsvArg.Chosen == 2 {
		err = pl.readCsv(cmd)

		if err == nil {
			fmt.Println("import datasource from csv is successfully")
		}

		return
	}

	if cmd.Arguments.CsvArg.Chosen == 1 {
		pl.repo.InitiateDefaultUser()

		fmt.Println("import default datasource successfully")

		return
	}

	err = utils.ErrorCommand

	return
}

func (pl *readCsv) readCsv(cmd *schemas.Command) (err error) {
	file, err := os.Open(cmd.Arguments.CsvArg.Path)

	// Checks for the error
	if err != nil {
		return
	}

	// Closes the file
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		return
	}

	if len(records) < 2 {
		err = utils.ErrorFileEmpty
		return
	}

	// validate the records
	var wg sync.WaitGroup
	errorCh := make(chan error)
	for _, record := range records[1:] {
		wg.Add(1)
		go pl.processRecord(&wg, record, errorCh)
	}

	go func() {
		wg.Wait()
		close(errorCh)
	}()

	for e := range errorCh {
		if e != nil {
			err = e
			break
		}
	}

	return
}

func (pl *readCsv) processRecord(wg *sync.WaitGroup, record []string, err chan error) {
	defer wg.Done()

	if record[0] == "" || len(record) < 1 {
		err <- utils.ErrorFileEmpty
		return
	}

	data := strings.Split(record[0], ";")
	if len(data) != 4 {
		err <- utils.ErrorValuesRecordInvalid
		return
	}

	balance, _ := strconv.Atoi(data[2])
	user := datasource.User{
		Id:            -1,
		Name:          data[0],
		Pin:           data[1],
		Balance:       int64(balance),
		AccountNumber: data[3],
	}

	// check user to datasource
	existingUser, _ := pl.repo.GetUserByAccountNumber(user.AccountNumber)
	if existingUser != nil {
		err <- utils.ErrorAccountNumberAlreadyExist
		return
	}

	// insert to datasource
	errInsert := pl.repo.InsertUser(user)
	if errInsert != nil {
		err <- utils.ErrorOccurs
		return
	}
}
