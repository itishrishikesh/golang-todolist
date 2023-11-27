package main

import "fmt"

// Simple struct
type furniture struct {
	cost     int
	serial   string
	width    int
	height   int
	breadth  int
	material string
	color    string
}

// Example of nested struct
type bed struct {
	details             furniture
	storageCompartments int
	isMattressIncluded  bool
}

// Example of struct with nested anonymous struct
type sofa struct {
	details     furniture
	accessories struct {
		isTableIncluded  bool
		isCusionIncluded bool
	}
}

// Example of embedded struct
type table struct {
	furniture
	legs       int
	isStanding int
}

// Example of a method
func (b *bed) addMattressDiscount() int {
	if b.isMattressIncluded {
		return 0
	}
	b.details.cost = b.details.cost - 10
	return b.details.cost
}

func main() {
	// Initializing an empty struct
	myBed := bed{}
	// Accessing a property
	myBed.details.cost = 10000
	fmt.Println("Cost of my bed is", myBed.details.cost)
	// --------------------------------------------------------- //
	// Anonymous struct (structs with no name)
	myChair := struct {
		details furniture
		typ     string
	}{
		typ: "office",
	}
	fmt.Println("My chair is for", myChair.typ)
	// --------------------------------------------------------- //
	// Accessing values from a struct which has an embedded struct
	myTable := table{}
	myTable.cost = 10000
	fmt.Println("Cost of my table is", myTable.cost)

	// --------------------------------------------------------- //
	// Method call
	myBed.details.cost = 10000
	myBed.isMattressIncluded = false
	fmt.Println("After adding no mattress discount", myBed.addMattressDiscount())
}
