package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var (
	mg MongoInstance
)

func DbConnect() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found (this is fine in production)")
	}

	dbName := os.Getenv("MONGO_DB_DATABASE")
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")
	port := os.Getenv("MONGO_DB_PORT")
	host := os.Getenv("MONGO_DB_HOST")

	var mongoURL = "mongodb://" + username + ":" + password + "@" + host + ":" + port + "/" + dbName + "?authMechanism=SCRAM-SHA-256&authSource=admin"

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
