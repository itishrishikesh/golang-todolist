package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func doSomething(something string, x, y int, arithmatic func(int, int) int) int {
	fmt.Println(something, ":")
	return arithmatic(arithmatic(x, y), y)
}

func main() {
	fmt.Println(doSomething("Random Arithmetics with add func", 100, 10, add))
	fmt.Println(doSomething("Random Arithmetics with sub func", 100, 10, sub))
}
