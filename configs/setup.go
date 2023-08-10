package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Declare a global variable to store the MongoDB client
var Client *mongo.Client

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:1234@localhost:27017/")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the global variable to the MongoDB client
	Client = client
}
