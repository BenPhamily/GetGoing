package controller

import (
	"net/http"

	"github.com/BenPhamily/GetGoing/todoList/model"
)

func find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// take some data and save it
			if err := model.Find(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}
