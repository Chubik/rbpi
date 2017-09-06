package main

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	motion := gpioreg.ByName("17")
	fmt.Printf("%s: %s\n", motion, motion.Function())
	led := gpioreg.ByName("26")
	fmt.Printf("%s: %s\n", led, led.Function())

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
