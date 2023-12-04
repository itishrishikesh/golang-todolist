package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type battery struct {
	power map[string]int
	mux   *sync.Mutex
}

func (b *battery) light(burnPower int) {
	b.mux.Lock()
	defer b.mux.Unlock()
	time.Sleep(time.Millisecond * time.Duration(burnPower))
	b.power["light"] -= burnPower
	fmt.Println("Power: ", b.power)
	wg.Done()
}

func (b *battery) fan(burnPower int) {
	b.mux.Lock()
	defer b.mux.Unlock()
	fmt.Println("Power: ", b.power)
	b.power["fan"] -= burnPower
	wg.Done()
}

func main() {
	myBattery := battery{
		power: map[string]int{
			"light": 100,
			"fan":   200,
		},
		mux: &sync.Mutex{},
	}

	wg.Add(1)
	go myBattery.light(2)
	wg.Add(1)
	go myBattery.light(20)
	wg.Add(1)
	go myBattery.light(33)
	wg.Add(1)
	go myBattery.light(233)
	wg.Add(1)
	go myBattery.light(123)
	wg.Add(1)
	go myBattery.light(123)
	wg.Add(1)
	go myBattery.light(11)
	// wg.Add(1)
	// go myBattery.fan(1)

	wg.Wait()
}
