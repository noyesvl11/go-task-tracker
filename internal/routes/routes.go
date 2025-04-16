package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/delivery"
	"rest-project/internal/repository"

	service "rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine) {
	// Auth routes
	authGroup := r.Group("api/v1/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
	}

	// Task routes
	taskRepo := repository.NewTaskRepository(db.DB)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := delivery.NewTaskHandler(taskService)

	taskGroup := r.Group("api/v1/tasks")
	{
		taskGroup.GET("/", taskHandler.GetAllTasks)
		taskGroup.GET("/:id", taskHandler.GetTask)
		taskGroup.POST("/", taskHandler.CreateTask)
		taskGroup.PUT("/:id", taskHandler.UpdateTask)
		taskGroup.DELETE("/:id", taskHandler.DeleteTask)
	}
}
