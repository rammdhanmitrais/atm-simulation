package utils

import "errors"

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
