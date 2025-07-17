package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB establishes a connection to MongoDB and returns the client and a context.
func ConnectDB(uri string) (*mongo.Client, context.Context, context.CancelFunc) {
	// Use a context with a timeout to prevent the application from hanging indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Set client options and connect.
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the primary to verify that the connection is alive.
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB!")
	return client, ctx, cancel
}
