package main

import (
	"context"
	"log"
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	infrastructure "task-manager/Infrastructure"
	repositories "task-manager/Repositories"
	usecases "task-manager/Usecases"
	"task-manager/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	client, err := connectToDatabase(cfg.Database.URI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	database := client.Database(cfg.Database.Database)

	// Initialize infrastructure services
	passwordService := infrastructure.NewPasswordService()
	authService := infrastructure.NewAuthService()

	// Initialize repositories
	taskRepo := repositories.NewTaskRepository(database.Collection("tasks"))
	userRepo := repositories.NewUserRepository(database.Collection("users"), passwordService)

	// Initialize use cases
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(userRepo, passwordService, authService)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)

	// Setup router with middleware
	r := routers.SetupRouter(taskController, userController, authService)

	// Start server
	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Starting server on %s", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func connectToDatabase(uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}
