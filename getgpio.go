package main

import (
	"fmt"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	host.Init()
	fmt.Println(gpioreg.All())
}
