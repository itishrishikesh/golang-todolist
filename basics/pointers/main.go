package main

import "fmt"

// Pointer is just a variable that stores the memory address

func main() {

	value := 10

	// ampersand is used to create reference a value
	pointer := &value

	fmt.Println("10 is stored at", pointer)

	// asterisk is a derefernce operator
	deferencedValue := *pointer

	fmt.Println(pointer, "is deferenced as value", deferencedValue)

	modifyOriginalValue(pointer)

	fmt.Println("The value of variable after changing it through a func is", value)
}

func modifyOriginalValue(value *int) {
	// *value = *value + 10
	myValue := *value
	myValue = myValue + 10
	*value = myValue
}
