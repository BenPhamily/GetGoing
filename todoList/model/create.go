package model

import (
	"context"
	"fmt"

	"github.com/BenPhamily/GetGoing/todoList/views"
)

func CreateTodo() error {
	collection := con.Database("TodoList").Collection("Todo")
	data := views.Todo{
		Name: "Ben",
		Todo: "this project",
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	id := res.InsertedID
	fmt.Println(id)
	return nil
}
