package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "https://github.com/BenPhamily/GetGoing/todoList/structs"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world go"))
	})

	http.ListenAndServe(":3000", nil)
}
