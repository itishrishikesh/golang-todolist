package main

import (
	"fmt"
	"strings"
)

const ZERO int = 0
const NUMBER_OF_DAY_IN_WEEK int = 7

var days []string = []string{"sunday", "monday", "tuesday", "wednesday", "thurday", "friday", "saturday"}

var WORK []string = []string{"monday", "tuesday", "wednesday", "thurday", "friday"}
var HOLIDAY []string = []string{"sunday", "saturday"}

func main() {
	// simple loop.
	for i := ZERO; i < NUMBER_OF_DAY_IN_WEEK; i++ {
		fmt.Printf(" -> %s", days[i])
	}
	// written like while loop of other languages
	fmt.Println("\nworking days:")
	current_day := "monday"
	current_day_index := 0
	for isWorkingDay(current_day) {
		current_day = days[current_day_index]
		fmt.Print(" ->-> ", current_day)
		current_day_index += 1
	}
}

func isWorkingDay(day string) bool {
	for _, d := range WORK {
		if strings.EqualFold(day, d) {
			return true
		}
	}
	return false
}
