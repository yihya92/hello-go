package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID    int64  `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Age   int    `json:"age" bson:"age"`
}

func InsertUser(user User) error {
	collection := mongoClient.Database(db).Collection(collName)

	// Count documents
	count, _ := collection.CountDocuments(context.TODO(), bson.M{})

	user.ID = count + 1

	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func InsertManyUsers(users []User) error {
	collection := mongoClient.Database(db).Collection(collName)

	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var documents []interface{}

	for i := range users {
		users[i].ID = count + int64(i) + 1
		documents = append(documents, users[i])
	}

	_, err = collection.InsertMany(context.TODO(), documents)
	return err
}

func GetUserByID(id int) (User, error) {
	collection := mongoClient.Database(db).Collection(collName)

	var user User

	err := collection.FindOne(
		context.TODO(),
		bson.M{"id": id},
	).Decode(&user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(id int64, user User) error {
	collection := mongoClient.Database(db).Collection(collName)

	filter := bson.M{"id": id}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func DeleteUser(userID int64) error {
	collection := mongoClient.Database(db).Collection(collName)

	filter := bson.M{"id": userID}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}

	fmt.Println("Deleted record count:", result.DeletedCount)
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
