package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	// switch case
	// select for channel async
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("Received")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(0)
		}
	}
}

func main() {
	c := make(chan int, 2)
	quits := make(chan struct{})
	go Select(c, quits)

	go func() {
		c <- 1
		quits <- struct{}{}
	}()

	select {}
}
