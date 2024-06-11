package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database Instance
func DBInstance() *mongo.Client {
	MongoDBURI := os.Getenv("MongoDBURI")
	fmt.Println(MongoDBURI)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDBURI))
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🏳️🏳️🏳️ -------- Connected to MongoDB ------------ 🥳🥳🥳")
	return client
}

// OpenCollection is a function that makes a connection with a collection in the database
func OpenCollection(Client *mongo.Client, collectionName string) *mongo.Collection {
	log.Printf("Created Collection: %v", collectionName)
	var collection *mongo.Collection = Client.Database("auth-server").Collection(collectionName)

	return collection
}