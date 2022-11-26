package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/BenPhamily/GetGoing/todoList/controller"
	"github.com/BenPhamily/GetGoing/todoList/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Disconnect(context.Background())
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
