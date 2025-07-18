package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.TaskController, ac *controllers.AuthController) *gin.Engine {
	r := gin.Default()
	r.POST("/register", ac.Register)
	r.POST("/login", ac.Login)

	r.GET("/tasks", tc.GetAllTasks)
	r.GET("/tasks/:id", tc.GetTaskByID)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		adminRoutes := authRoutes.Group("/")
		adminRoutes.Use(middleware.AdminOnly())
		{
			adminRoutes.POST("/tasks/", tc.CreateTask)
			adminRoutes.PUT("/tasks/:id", tc.UpdateTask)
			adminRoutes.DELETE("/tasks/:id", tc.DeleteTask)
		}
	}
	return r
}
