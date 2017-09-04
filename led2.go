package main

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	//	firmataAdaptor := firmata.NewAdaptor("COM111")
	r := raspi.NewAdaptor()
	laser := gpio.NewLedDriver(r, "13")

	work := func() {
		laser.On()

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{r},
		[]gobot.Device{laser},

		work,
	)

	robot.Start()

}
