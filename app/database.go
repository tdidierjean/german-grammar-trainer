package grammarexercise

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "test"
const userExerciseTypeCollection = "UserExerciseType"

// Connection client wrapper
type Connection struct {
	client *mongo.Client
}

// UserExerciseType schema for storing a user's exercise type choice
type UserExerciseType struct {
	UserID       int
	ExerciseType string
}

// Connect attempts to connect to the MongoDB instance, returns true if successful
func (c *Connection) Connect() bool {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODBCONNSTR"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
		return false
	}

	c.client = client

	// Check the connection
	err = c.client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("Connected to MongoDB!")
	return true
}

// GetUserExerciseType return the stored type of exercise for this user if it exists
func (c *Connection) GetUserExerciseType(userID int) string {
	filter := bson.D{{"userid", userID}}

	collection := c.client.Database(databaseName).Collection(userExerciseTypeCollection)

	var found UserExerciseType
	err := collection.FindOne(context.TODO(), filter).Decode(&found)
	if err != nil {
		return ""
	}

	return found.ExerciseType
}

// UpdateUserExerciseType inserts or update the stored exercise type for the user
func (c *Connection) UpdateUserExerciseType(userID int, exerciseType string) {
	filter := bson.D{{"userid", userID}}

	update := bson.D{
		{"$set", bson.D{
			{"exercisetype", exerciseType},
		}},
	}

	collection := c.client.Database(databaseName).Collection(userExerciseTypeCollection)
	_, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		log.Fatal(err)
	}
}
