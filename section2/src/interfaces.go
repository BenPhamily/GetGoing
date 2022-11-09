package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo")
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"C369"}
	l.Drive()
	c.Drive()
	car1 := NewModel("Biden")
	car1.Stop()
}
