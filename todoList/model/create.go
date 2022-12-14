package model

import (
	"context"
	"fmt"

	"github.com/BenPhamily/GetGoing/todoList/views"
)

func CreateTodo(name, todo string) error {
	collection := con.Database("TodoList").Collection("Todo")
	data := views.Todo{
		Name: name,
		Todo: todo,
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	id := res.InsertedID
	fmt.Println(id)
	return nil
}
