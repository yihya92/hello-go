package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("mydb").Collection("users")
	createEmailIndex(collection, ctx)
	id := insertUser(collection, ctx)
	findOneUser(collection, ctx, id)
	findUsers(collection, ctx)
	//updateUserAge(collection, ctx, id)
	//deleteUser(collection, ctx, id)
}
