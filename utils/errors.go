package utils

import (
	"errors"
	"fmt"
)

var ErrorInvalidAccount = errors.New("invalid account")
var ErrorInvalidAmount = errors.New("invalid amount")
var ErrorUserLogged = errors.New("no user logged")
var ErrorInvalidAccountPin = errors.New("invalid account number/pin")
var ErrorMaximumAmountTransfer = errors.New("maximum amount to transfer is $1000")
var ErrorMinimumAmountTransfer = errors.New("minimum amount to transfer is $1")
var ErrorReferenceNumber = errors.New("invalid reference number")
var ErrorAccountNumberNotNumber = errors.New("account number should only contains numbers")
var ErrorAccountNumberDigit = errors.New("account number should have 6 digits length")
var ErrorPINNotNumber = errors.New("PIN should only contains numbers")
var ErrorPINDigit = errors.New("PIN should have 6 digits length")
var ErrorExit = errors.New("exit")
var ErrorCommand = errors.New("command invalid")
var ErrorValuesRecordInvalid = errors.New("record has invalid values")
var ErrorAccountNumberAlreadyExist = errors.New("account number already exist")
var ErrorOccurs = errors.New("there are problems happen")

func SetErrorInsufficient(curr string, balance int64) error {
	return fmt.Errorf("insufficient balance %s%d", curr, balance)
}

func SetErrorDuplicateAccountNumber(accountNumber string) error {
	return fmt.Errorf("there can't be 2 different accounts with the same account number %s", accountNumber)
}

func SetErrorDuplicateRecord(record string) error {
	return fmt.Errorf("there can't be duplicated records %s", record)
}
