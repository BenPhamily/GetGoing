package main

import (
	"fmt"
)

type Car struct {
	Model string
}

// Buffered Channels
func main() {
	// Set buffered channel size 3
	c := make(chan *Car, 3)

	go func() {
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		// Close the channel
		close(c)
	}()

	for i := range c {
		fmt.Println(i.Model)
	}

}
