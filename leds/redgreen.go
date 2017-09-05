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
		gpioreg.ByName("6").Out(l)
		time.Sleep(500 * time.Millisecond)
	}
	// gpioreg.ByName("13").Out(gpio.High)
	// gpioreg.ByName("6").Out(gpio.High)

}
