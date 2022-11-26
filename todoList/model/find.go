package model

import (
	"context"
	"fmt"
	"log"

	"github.com/BenPhamily/GetGoing/todoList/views"
	"go.mongodb.org/mongo-driver/bson"
)

func Find() error {
	collection := con.Database("TodoList").Collection("Todo")
	// Several methods return a cursor, which can be used like this:

	cur, err := collection.Find(context.Background(), bson.M{"name": "Ben"})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		// To decode into a struct, use cursor.Decode()
		result := views.Todo{}
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result...

		// To get the raw bson bytes use cursor.Current
		raw := cur.Current
		// do something with raw...
		fmt.Println(raw)
	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}
