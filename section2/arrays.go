package main

import "fmt"

func todo() {
	// var arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}

	arr2 = append(arr2, "is ben")
	fmt.Println(arr, arr2)

}

func main() {
	todo()

}
