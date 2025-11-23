package config

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToDatabase() (*mongo.Database, error) {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatal("Mongo URI is not set")
	}

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(opts)

	if err != nil {
		return nil, err
	}

	dbName := os.Getenv("MONGO_DB")

	if dbName == "" {
		log.Fatal("Mongo DB Name is not set")
	}

	db := client.Database(dbName)
	return db, nil
}
