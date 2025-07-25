package main

import (
	"context"
	"log"
	"os"
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	"task-manager/Repositories"
	"task-manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Ping the primary
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Successfully connected and pinged MongoDB.")

	database := client.Database("task_manager_db")

	taskRepo := repositories.NewTaskRepository(database.Collection("tasks"))
	userRepo := repositories.NewUserRepository(database.Collection("users"))

	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)

	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)

	r := routers.SetupRouter(taskController, userController)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
