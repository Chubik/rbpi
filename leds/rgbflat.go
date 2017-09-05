package main

import (
	"sync"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	var wg sync.WaitGroup
	host.Init()
	// for l := gpio.Low; ; l = !l {
	// 	gpioreg.ByName("13").Out(l)
	// 	time.Sleep(500 * time.Millisecond)
	// 	gpioreg.ByName("6").Out(l)
	// 	time.Sleep(500 * time.Millisecond)
	// 	gpioreg.ByName("5").Out(l)
	// 	time.Sleep(500 * time.Millisecond)
	// }
	wg.Add(3)
	go ledOUT("13")
	go ledOUT("6")
	go ledOUT("5")
	// go func() {
	// 	for i := range lights {
	// 		fmt.Println(i)
	// 	}
	// }()

	wg.Wait()
}

func ledOUT(name string) {
	for l := gpio.Low; ; l = !l {
		gpioreg.ByName(name).Out(l)
		time.Sleep(500 * time.Millisecond)
	}
}
