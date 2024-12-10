package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var (
	mg *MongoInstance
)

const dbName = "service-order"
const mongoURL = "mongodb://admin:password@localhost:27017/" + dbName

func DbConnect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg.Client = client
	mg.Db = db

	return nil
}

func GetMongoInstance() *MongoInstance {
	return mg
}
