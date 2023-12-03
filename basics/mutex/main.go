package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type battery struct {
	power int
	mux   *sync.Mutex
}

func (b *battery) light() {
	b.mux.Lock()
	time.Sleep(time.Second * time.Duration(b.power))
	b.power -= 1
	fmt.Println("Power: ", b.power)
	wg.Done()
	b.mux.Unlock()
}

// func (b *battery) fan() {
// 	b.mux.Lock()
// 	time.Sleep(time.Second * time.Duration(b.power))
// 	b.power -= 30
// 	fmt.Println("Power: ", b.power)
// 	wg.Done()
// 	b.mux.Unlock()
// }

func main() {
	myBattery := battery{
		power: 300,
		mux:   &sync.Mutex{},
	}

	wg.Add(1)
	go myBattery.light()
	wg.Add(1)
	go myBattery.light()
	wg.Add(1)
	go myBattery.light()
	// wg.Add(1)
	// go myBattery.fan()
	// go myBattery.light()
	// go myBattery.light()

	wg.Wait()
}
