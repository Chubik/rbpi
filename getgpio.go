package main

import (
	"fmt"

	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

func main() {
	host.Init()
	for _, gi := range gpioreg.All() {
		fmt.Println(gpioreg.ByName(gi).Name)
	}

}
