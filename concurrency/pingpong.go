package main

import (
	"fmt"
	"time"
)

// pinger prints ping and waits for pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// ponger prints pong and waits for ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int, 0)
	pong := make(chan int, 0)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// Initialise main go routine by sending into the ping channel
	ping <- 1

	for {
		// Block main thread until interrupt
		time.Sleep(time.Second)
	}
}
