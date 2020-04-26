package fyne_utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var regexpFloatNumber  = regexp.MustCompile(`\d*\.*\d*`)

func NumericInDiapason(input *string,from float64, to float64, defaultValue float64) error {
	number,err := strconv.ParseFloat(*input,64);
	var resultError error

	if err != nil {
		number = defaultValue
		resultError =  err
	}

	if number < from || number > to {
		number = defaultValue
		resultError = errors.New(fmt.Sprintf("number beyond the borders. should be between %.2f and %.2f",from, to))
	}

	*input = fmt.Sprint(number)
	return resultError
}

func Numeric(input *string){
	*input = regexpFloatNumber.FindString(*input)
}