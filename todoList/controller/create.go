package controller

import (
	"net/http"

	"github.com/BenPhamily/GetGoing/todoList/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// take some data and save it
			if err := model.CreateTodo(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}