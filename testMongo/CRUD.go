package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertUser(collection *mongo.Collection, ctx context.Context) primitive.ObjectID {
	user := User{
		Name:  "Kamal",
		Age:   33,
		Email: "kamal.lakkis@gmail.com",
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		panic(err)
	}

	id := result.InsertedID.(primitive.ObjectID)
	fmt.Println("Inserted ID:", id.Hex())

	return id
}

func findOneUser(collection *mongo.Collection, ctx context.Context, id primitive.ObjectID) {
	filter := struct {
		ID primitive.ObjectID `bson:"_id"`
	}{
		ID: id,
	}

	var user User

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found")
			return
		}
		panic(err)
	}

	fmt.Printf("Found user: %+v\n", user)
}

func findUsers(collection *mongo.Collection, ctx context.Context) {
	filter := struct {
		Age struct {
			GT int `bson:"$gt"`
		} `bson:"age"`
	}{
		Age: struct {
			GT int `bson:"$gt"`
		}{
			GT: 18,
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			panic(err)
		}
		fmt.Println("User:", user)
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}
}

func updateUserAge(collection *mongo.Collection, ctx context.Context, id primitive.ObjectID) {
	filter := struct {
		ID primitive.ObjectID `bson:"_id"`
	}{
		ID: id,
	}

	update := UpdateAge{}
	update.Set.Age = 30

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Println("Matched:", result.MatchedCount)
	fmt.Println("Modified:", result.ModifiedCount)
}

func deleteUser(collection *mongo.Collection, ctx context.Context, id primitive.ObjectID) {
	filter := struct {
		ID primitive.ObjectID `bson:"_id"`
	}{
		ID: id,
	}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted count:", result.DeletedCount)
}

func createEmailIndex(collection *mongo.Collection, ctx context.Context) {
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "email", Value: 1}, // 1 = ascending
		},
		Options: options.Index().
			SetUnique(true).
			SetName("unique_email"),
	}

	name, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created index:", name)
}
