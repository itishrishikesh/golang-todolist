package main

import "fmt"

// Simple interface
type ninja interface {
	rank() string
	jutsu() string
}

type hokage struct {
	rank         int
	primaryJutsu string
}

func main() {
	fmt.Println()
}
