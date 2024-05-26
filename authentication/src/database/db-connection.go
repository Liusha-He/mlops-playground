package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Client {
	uri := os.Getenv("MONGODB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to db")
	return client
}

var Client *mongo.Client = DB()

func NewCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("").Collection(collectionName)
	return collection
}
