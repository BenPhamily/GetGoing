package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println()
	Anything(2.44)
	Anything("Ben")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "BenPhamily"
	mymap["age"] = 10
	fmt.Println(mymap)
}
