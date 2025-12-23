package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"exerciciomongdb/internal/http/routes"
	mongoplaform "exerciciomongdb/internal/plataform/mongo"
	"exerciciomongdb/internal/user"
)

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return value
}

func main() {
	_ = godotenv.Load()

	mongoURI := mustGetEnv("MONGO_URI")
	mongoDB := mustGetEnv("MONGO_DB")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongoplaform.Connect(ctx, mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func(c *mongoDriver.Client) {
		_ = c.Disconnect(context.Background())
	}(client)

	collection := mongoplaform.Collection(client, mongoDB, "users")
	repo := user.NewMongoRepository(collection)
	h := user.NewHandlers(repo)

	router := routes.NewRouter(h)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
