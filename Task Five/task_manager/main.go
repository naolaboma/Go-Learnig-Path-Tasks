package main

import (
	"log"
	"os"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/db"
	"task_manager/router"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	client, c, cancel := db.ConnectDB(mongoURI)
	defer cancel()
	defer func() {
		if err := client.Disconnect(c); err != nil {
			log.Fatalf("failed to disconnect from mongodb: %v", err)

		}
	}()

	database := client.Database("task_manager_dn")
	taskCollection := database.Collection("tasks")

	taskService := data.NewTaskService(taskCollection)
	taskController := controllers.NewTaskController(taskService)

	r := router.SetupRouter(taskController)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
