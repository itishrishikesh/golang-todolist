package main

import (
	"fmt"
	"sync"
)

func printWords(i int, word string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%d: %s\n", i, word)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"hello", "world", "earth", "sun", "god",
	}
	for i, word := range words {
		go printWords(i+1, word, &wg)
	}
	wg.Add(len(words))
	wg.Wait()
}
