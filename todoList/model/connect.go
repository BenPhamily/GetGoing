package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BenPhamily/GetGoing/todoList/controller"
)

func Connect() *sql.DB {
	db, err := sql.Open(
		"mysql", controller.GoDotEnvVariable("DSN"),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	return db
}
