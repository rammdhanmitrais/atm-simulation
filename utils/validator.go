package utils

import (
	"regexp"
)

var PatternNumber string = "^[0-9]+$"

func ValidateIsContainNumberOnly(number string) bool {
	var digitCheck = regexp.MustCompile(PatternNumber)
	return digitCheck.MatchString(number)
}
