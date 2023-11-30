package main

import (
	"fmt"
)

func main() {
	var thisIsAnArray [7]string = [7]string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println("Values in the array:", thisIsAnArray)

	var thisIsASlice []string = thisIsAnArray[0:5]
	fmt.Println("Values in the slice:", thisIsASlice)

	// You can append a slice.
	appendedSlice := append(thisIsASlice, "friday", "saturday")
	fmt.Println("Values in appended slice:", appendedSlice)

	// You cannot append an array. The append method only accepts slice.

	// Changing the slice will lead to updation of the underlying array.
	thisIsASlice[0] = "ravi"
	fmt.Println("Values in the slice:", thisIsASlice)
	fmt.Println("Values in the underlying array:", thisIsAnArray)
	thisIsASlice[0] = "sunday"

	// While passing the value of an array to a function it'll send copy of the array. As its passed by value.
	changeValueOfAnArray(thisIsAnArray)
	fmt.Println("Value of array in main func:", thisIsAnArray)

	// But when you pass a slice, it's only a reference to the original array. And the array values are changed.
	changeValueOfASlice(thisIsASlice)
	fmt.Println("Value of slice in main func:", thisIsAnArray)
	fmt.Println("Value of array in main func:", thisIsAnArray)
}

func changeValueOfAnArray(days [7]string) {
	days[0] = "ravi"
	fmt.Println("Value of the array in a differnt func:", days)
}

func changeValueOfASlice(days []string) {
	days[0] = "ravi"
	fmt.Println("Value of slice in a differnt func:", days)
}
