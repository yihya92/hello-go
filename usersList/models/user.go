package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
	Age   int                `json:"age" bson:"age"`
}

func InsertUser(user User) error {
	collection := mongoClient.Database(db).Collection(collName)
	inserted, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a record with id: ", inserted.InsertedID)
	return err
}

func InsertManyUsers(users []User) error {
	newUsers := make([]interface{}, len(users))
	for i, user := range users {
		newUsers[i] = user
	}
	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.InsertMany(context.TODO(), newUsers)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	return err
}

func FindUser(username string) User {
	var result User
	filter := bson.D{{Key: "user", Value: username}}
	collection := mongoClient.Database(db).Collection(collName)
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func FindAllUsers(username string) []User {
	var results []User
	filter := bson.D{{Key: "user", Value: username}}
	collection := mongoClient.Database(db).Collection(collName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func UpdateUser(userID string, user User) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": user.Name, "age": user.Age}}
	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("updated a record: ", result)
	return nil
}

func DeleteUser(userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	collection := mongoClient.Database(db).Collection(collName)
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Deleted a record: ", result)
	return nil
}

func ListAll() []User {

	var results []User

	collection := mongoClient.Database(db).Collection(collName)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func DeleteAll() error {
	collection := mongoClient.Database(db).Collection(collName)
	delResult, err := collection.DeleteMany(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		return err
	}

	fmt.Println("Records deleted: ", delResult.DeletedCount)
	return err
}
