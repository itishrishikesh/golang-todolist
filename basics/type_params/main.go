package main

import "fmt"

func splitAnySlice[T any](a []T) ([]T, []T) {
	mid := len(a) / 2
	return a[mid:], a[:mid]
}

func main() {

	powers := []int{2, 4, 16, 32}

	bottles := []string{"copper", "plastic", "steel"}

	fmt.Println(splitAnySlice[int](powers))
	fmt.Println(splitAnySlice[string](bottles))
}
