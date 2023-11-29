package main

import (
	"fmt"
	"strings"
)

var WORK []string = []string{"MON", "TUE", "WED", "THUS", "FRI"}
var HOLIDAY []string = []string{"SAT", "SUN"}

// A simple function that return an error
func shouldICodeToday(day string) (goodDayToCode bool, err error) {
	for _, val := range WORK {
		if strings.EqualFold(val, day) {
			return true, err
		}
	}
	err = fmt.Errorf("Sorry can't work on days except for: %s", strings.Join(WORK[:], ","))
	return false, err
}

// Struct that implements Error
type coderError struct {
	workingDays []string
	holidays    []string
}

func (c coderError) Error() string {
	return fmt.Sprintf("Coder doesn't like to work on holidays: [%s]\n"+
		"Coder can't be forced to work on working days: [%s] and Sunday\n",
		strings.Join(c.holidays, ","), strings.Join(c.workingDays, ","))
}

func forceCoderToCode(day string) (coderIsForced bool, err error) {
	if strings.EqualFold(day, "Sat") {
		return true, nil
	}
	return false, coderError{
		holidays:    HOLIDAY,
		workingDays: WORK,
	}
}

func main() {
	// result, err := shouldICodeToday("JAN")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Coder works = ", result)
	result, err := forceCoderToCode("SUN")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Program checked if coder works? and the answer in boolean is", result)
}
