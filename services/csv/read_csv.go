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
			fmt.Println("generate csv datasource successfully")
		}

		return
	}

	if cmd.Arguments.CsvArg.Chosen == 1 {
		pl.repo.InitiateDefaultUser()

		fmt.Println("generate default datasource successfully")

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

	accountNumberRead := make(map[string]bool)
	readyRecord := []datasource.User{}

	// validate the records
	for _, eachrecord := range records {
		if eachrecord[0] == "" || len(eachrecord) < 1 {
			err = utils.ErrorValuesRecordInvalid
			return
		}

		record := strings.Split(eachrecord[0], ";")
		if len(record) != 4 {
			err = utils.ErrorValuesRecordInvalid
			return
		}

		// check account number duplicate or no
		if accountNumberRead[record[3]] {
			err = utils.SetErrorDuplicateAccountNumber(record[3])
			return
		}

		balance, _ := strconv.Atoi(record[2])
		user := datasource.User{
			Id:            -1,
			Name:          record[0],
			Pin:           record[1],
			Balance:       int64(balance),
			AccountNumber: record[3],
		}

		accountNumberRead[record[3]] = true
		readyRecord = append(readyRecord, user)
	}

	// validate record to datasource
	for _, user := range readyRecord {
		// check user to datasource
		existingUser, _ := pl.repo.GetUserByAccountNumber(user.AccountNumber)
		if existingUser != nil {
			err = utils.ErrorAccountNumberAlreadyExist
			return
		}

		// insert to datasource
		err = pl.repo.InsertUser(user)
		if err != nil {
			err = utils.ErrorOccurs
			return
		}
	}

	return
}
