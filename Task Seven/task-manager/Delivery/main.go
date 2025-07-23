// Delivery/main.go
package main

import (
	"context"
	"log"
	"os"
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	"task-manager/Domain"
	"task-manager/Infrastructure"
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
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	database := client.Database("task_manager_db")
	
	// Initialize repositories
	taskRepo := repositories.NewTaskRepository(database.Collection("tasks"))
	userRepo := repositories.NewUserRepository(database.Collection("users"))

	// Initialize use cases
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)

	// Setup router
	r := routers.SetupRouter(taskController, userController)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}