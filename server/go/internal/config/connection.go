package config

import (
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToDatabase() *mongo.Database {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		slog.Error("Mongo URI is not set")
		return nil
	}

	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(opts)

	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	dbName := os.Getenv("MONGO_DB")

	if dbName == "" {
		slog.Error("Mongo DB Name is not set")
		return nil
	}

	return client.Database(dbName)
}
