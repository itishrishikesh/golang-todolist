package main

import (
	"fmt"
	"strings"
)

// Simple Function
func add(n1, n2 int) int {
	return n1 + n2
}

// Parameters are passed by value
func increment(v int) int {
	v = v + 1
	return v
}

// Returning multiple values
func getMinutesAndHour(timer string) (min, hour string) {
	var arr []string = strings.Split(timer, ":")
	min = arr[0]
	hour = arr[1]
	return min, hour
}

func main() {
	// A simple example of a simple function
	fmt.Println("Adding 1 and 2: ", add(1, 2))

	// ----------------------------------------------- //

	// The value of x is passed by value
	var x int = 1
	fmt.Println("Original value of x: ", x)
	// Here we increment the value of x once.
	increment(x)
	fmt.Println("I'm trying to increment x: ", x)
	// Here too we increment the value of x once.
	// But the variable "x" is passed by value to function "increment".
	// Thus, the value that was set by the previous call doesn't change "x" here.
	increment(x)
	fmt.Println("I'm trying to increment x: ", x)
	// Same as above. Variable "x" doesn't get modified.
	increment(x)
	fmt.Println("I'm trying to increment x: ", x)

	// ----------------------------------------------- //

	// The above code can be modified to increment variable "x" if each time we're calling "increment"
	// we assign the return value of "increment" function to variable "x".
	fmt.Println("Original value of x: ", x)
	x = increment(x)
	fmt.Println("I incremented x:", x)
	x = increment(x)
	fmt.Println("I incremented x:", x)

	// ----------------------------------------------- //

	// Go doesn't allow unused variable.
	// If your function is returning multiple values, and you don't need one of them.
	// Then we can use the underscore character to ignore the value.
	hour, _ := getMinutesAndHour("12:08")
	fmt.Println("Hour is ", hour)
}
