package main

import "fmt"

type anime struct {
	name    string
	genre   string
	ratings int
}

var author map[string]anime

func main() {
	author = make(map[string]anime)
	author["john"] = anime{
		name:    "john doesn't do anime",
		genre:   "just random bull crap",
		ratings: 1000,
	}
	fmt.Println("Simple Map:", author)

	val, ok := author["john"]
	fmt.Println(val, ok)
}
