package main

import (
    "user-service/internal/handler"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.POST("/users", handler.CreateUser)
    router.GET("/users/:id", handler.GetUser)

    router.Run(":8081") // User service will run on port 8081
}
