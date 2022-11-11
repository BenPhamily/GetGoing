package main

import (
	"fmt"
	"time"
)

// Calling go routine as a function means it will be running in the background

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Hello")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super heavy")
	}
}

// If function exits then go routines will also exit
func main() {
	go heavy() // go routine
	go superHeavy()
	fmt.Println("Fin")
	select {} // infinite stall
}
