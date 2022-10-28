package main

import (
	"fmt"
)

// structures
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving..")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Name:    "chevy",
		Age:     1,
		ModelNo: 2,
	}
	// var c1 Car
	c2 := Car{"chevy", 2, 3}
	fmt.Println(c)
	c2.Print()
	c2.Drive()
	fmt.Println(c.GetName())
}
