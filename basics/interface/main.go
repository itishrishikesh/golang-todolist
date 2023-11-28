package main

import (
	"fmt"
	"strings"
)

// simple interface.
type furniture interface {
	changeMaterial(material string)
	changeCost(cost float64)
	getSeasonalDiscount(season string) int
}

type table struct {
	material string
	cost     float64
}

func (t table) changeMaterial(material string) {
	t.material = material
}

func (t table) changeCost(cost float64) {
	t.cost = cost
}

func (t table) getSeasonalDiscount(season string) (discount int) {
	if strings.EqualFold("winter", season) {
		discount = 10
	}
	discount = 20
	return discount
}

type trampoline struct {
	material string
	cost     float64
}

func (t trampoline) changeMaterial(material string) {
	t.material = material
}

func (t trampoline) changeCost(cost float64) {
	t.cost = cost
}

func (t trampoline) getSeasonalDiscount(season string) (discount int) {
	if strings.EqualFold("winter", season) {
		discount = 10
	}
	discount = 20
	return discount
}

func printFurnitureDetails(f furniture, s string) {
	// c, ok := f.(table)
	// if ok {
	// 	fmt.Println(s, "Material", c.material)
	// }
	// d, ok := f.(trampoline)
	// if ok {
	// 	fmt.Println(s, "Cost", d.cost)
	// }
	switch v := f.(type) {
	case table:
		fmt.Println(s, "Material", v.material)
	case trampoline:
		fmt.Println(s, "Cost", v.cost)
	default:
		fmt.Println("Something weird happened!")
	}
}

func main() {
	myTable := table{
		material: "wood",
		cost:     10000,
	}
	myTrampoline := trampoline{}

	printFurnitureDetails(myTable, "table")
	printFurnitureDetails(myTrampoline, "trampoline")

	// type assertions
	var f furniture = &table{
		material: "silicon",
		cost:     50000,
	}

	c, ok := f.(*table)
	fmt.Println("ok", ok)
	fmt.Println("c", c)

}
