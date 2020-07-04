package fyne_utils

import (
	"fmt"
	"regexp"
	"strconv"
)

var regexpFloatNumber = regexp.MustCompile(`\d*\.*\d*`)

func FilterNumericInDiapason(input string, from float64, to float64, defaultValue float64) (string, error) {
	number, err := strconv.ParseFloat(input, 64)
	var resultError error

	if err != nil {
		number = defaultValue
		resultError = err
	}

	if number < from || number > to {
		number = defaultValue
		resultError = fmt.Errorf("number beyond the borders. should be between %.2f and %.2f", from, to)
	}

	input = fmt.Sprint(number)
	return input, resultError
}

func FilterNumeric(input string) string {
	return regexpFloatNumber.FindString(input)
}
