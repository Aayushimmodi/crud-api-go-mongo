package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Load environment variables
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// ConnectDatabase initializes the MongoDB client
func ConnectDatabase() (*mongo.Client, error) {
	LoadEnv()
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB!")
	return client, nil
}
