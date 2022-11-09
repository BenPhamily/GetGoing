package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// if else, for, switch case, break continue, && ||

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	arr := []string{"firstname", "lastname", "middlename"}
	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "firstname"
	mymap["age"] = 20
	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v\n", k, v)
	}

}
