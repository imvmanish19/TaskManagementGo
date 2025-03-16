package main

import (
    "task-service/internal/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/tasks", handler.CreateTask)
    router.GET("/tasks", handler.GetTasks)
    router.PUT("/tasks/:id", handler.UpdateTask)
    router.DELETE("/tasks/:id", handler.DeleteTask)

    router.Run(":8080") // Task service will run on port 8080
}
