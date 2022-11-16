package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Multiplexer / Router
	mux := http.NewServeMux()

	// path matching the base route
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world"))
	})

	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		fmt.Println(r.Method)
		w.Write([]byte("Hello world go"))
	})

	http.ListenAndServe("localhost:3000", mux)
	// curl -i localhost:3000/		More Information
	// curl -X POST localhost:3000/	Set Method
}
