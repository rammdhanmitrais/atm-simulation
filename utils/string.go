package utils

import (
	"crypto/rand"
	"io"
	"regexp"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
var PatternNumber string = "^[0-9]+$"

func GenerateNDigitRandom(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	return string(b)
}

func ValidateIsContainNumberOnly(number string) bool {
	var digitCheck = regexp.MustCompile(PatternNumber)
	return digitCheck.MatchString(number)
}