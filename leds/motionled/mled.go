package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

var (
	gpioIN  string
	gpioOUT string
)

func main() {
	fmt.Println("Initialization...")
	gpioIN = "17"
	gpioOUT = "26"
	//TODO: periph states status
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded periph driver.")

	motion := gpioreg.ByName(gpioIN)

	led := gpioreg.ByName(gpioOUT)

	fmt.Printf("Initialized gpio IN %s , gpio OUT %s \n", gpioIN, gpioOUT)
	fmt.Printf("gpioIN %s: %s\n", motion, motion.Function())
	fmt.Printf("gpioOUT %s: %s\n", led, led.Function())

	if err := motion.In(gpio.PullDown, gpio.BothEdges); err != nil {
		log.Fatal(err)
	}

	for {
		motion.WaitForEdge(-1)
		if motion.Read() == false {
			led.Out(gpio.High)
		} else {
			led.Out(gpio.Low)
		}
		// fmt.Printf("-> %s\n", motion.Read())
	}
}
