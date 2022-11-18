package main

import (
	"log"
	"net/http"

	"github.com/BenPhamily/GetGoing/todoList/controller"
	"github.com/BenPhamily/GetGoing/todoList/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":3000", mux))
}
