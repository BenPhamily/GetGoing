package main

import (
	"fmt"
	"strings"
)

func main() {

	// Initialisation/Declaration
	m1 := "ben"
	m2 := "name"

	// Assignment
	m1 = "my name"

	fmt.Println(m1)
	fmt.Println(strings.Contains(m1, m2))

	// Concatenation, Split
	fmt.Println(strings.Split(m1, " "), m1+m2)
	fmt.Println(strings.ReplaceAll(m1, "m", "NO"))
}
