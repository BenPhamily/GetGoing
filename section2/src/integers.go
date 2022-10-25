package main

import "fmt"

func main() {

	var (
		m1 int32 = 2
		m2 int64 = 3
	)

	// Initialising variables
	m3 := 2
	m4 := 3

	fmt.Println(int64(m1) + m2)
	fmt.Println(m3 - m4)
}
