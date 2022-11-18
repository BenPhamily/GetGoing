package main

import (
	"net/http"

	"github.com/BenPhamily/GetGoing/todoList/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
