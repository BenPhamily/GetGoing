package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var con *mongo.Client

func Connect() *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(GoDotEnvVariable("DSN")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	con = client
	fmt.Println("Connected to the database")
	return client
}
