package handler

import (
    "net/http"
    "strconv"
    "user-service/internal/models"
    "github.com/gin-gonic/gin"
)

var users []models.User

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.ID = strconv.Itoa(len(users) + 1)
    users = append(users, user)

    c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
    id := c.Param("id")

    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
