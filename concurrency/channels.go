package main

import (
	"fmt"
)

// Unbuffered channels
func main() {
	c := make(chan int)
	// <name> chan <datatype>

	// send in a goroutine
	go func() {
		c <- 1
	}()

	// sniff
	val := <-c
	fmt.Println(val)

	// send in a goroutine
	go func() {
		c <- 2
	}()

	val = <-c
	fmt.Println(val)

	fmt.Println(c)
}
