package main

import (
	"time"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	host.Init()
	for l := gpio.Low; ; l = !l {
		gpioreg.ByName("13").Out(l)
		time.Sleep(500 * time.Millisecond)
		gpioreg.ByName("13").Out(gpio.High)
		time.Sleep(500 * time.Millisecond)
	}
}
