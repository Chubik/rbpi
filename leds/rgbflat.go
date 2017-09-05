package main

import (
	"fmt"
	"sync"
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

var (
	lights chan string
	wg     sync.WaitGroup
)

func main() {
	lights = make(chan string)
	host.Init()
	wg.Add(3)
	go ledOUT("13", 500)
	go ledOUT("6", 600)
	go ledOUT("5", 700)
	go func() {
		for i := range lights {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}

func ledOUT(name string, tm time.Duration) {
	for l := gpio.Low; ; l = !l {
		gpioreg.ByName(name).Out(l)
		time.Sleep(tm * time.Millisecond)
		lights <- name
	}
}
