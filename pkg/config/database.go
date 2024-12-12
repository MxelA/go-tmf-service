package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var (
	mg MongoInstance
)

const dbName = "service-order"
const mongoURL = "mongodb://admin:password@127.0.0.1:27017/" + dbName + "?authMechanism=SCRAM-SHA-256&authSource=admin"

func DbConnect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func GetMongoInstance() MongoInstance {
	return mg
}
