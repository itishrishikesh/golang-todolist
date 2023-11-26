package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type declaration
	var name string
	var age int
	var marks float64
	var pass bool

	name = "Ramesh"
	age = 45
	marks = 12.11
	pass = true

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Marks: ", marks)
	fmt.Println("Pass: ", pass)
	
	// infered declaration
	price := 2.4566
	bought := false
	product := "mobile"

	fmt.Println("Price: ", price)
	fmt.Println("Price type: ", reflect.TypeOf(price))
	fmt.Println("Bought: ", bought)
	fmt.Println("Bought type: ", reflect.TypeOf(bought))
	fmt.Println("Product: ", product)
	fmt.Println("Product type: ", reflect.TypeOf(product))

	// const
	const earth = "planet"
	fmt.Println("Earth is a ", earth)
	fmt.Println("Earth type: ", reflect.TypeOf(earth))
}