package db

import (
	"bookz/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	Client *mongo.Client
}

func (m *MongoDatabase) DBConnection() *mongo.Client {
	uri := config.GetEnv("MONGO_URL")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection Failed to Mongo Database")
		log.Fatal(err)
	}
	return client
}
