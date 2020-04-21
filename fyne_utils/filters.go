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
		resultError = errors.New("number beyond the borders")
	}

	*input = fmt.Sprint(number)
	return resultError
}

func Numeric(input *string){
	*input = regexpFloatNumber.FindString(*input)
}