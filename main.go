package main

import (
	"fmt"

	"github.com/AlejandroWaiz/singleton-pattern/singleton"
)

func main() {

	for i := 0; i < 5; i++ {

		go singleton.RescueThePrincess()

	}

	fmt.Scanln()

}
