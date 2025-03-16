package handler

import (
    "net/http"
    "strconv"
    "task-service/internal/models"
    "task-service/internal/grpc"
    "github.com/gin-gonic/gin"
)

var tasks []models.Task

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task.ID = strconv.Itoa(len(tasks) + 1)
    tasks = append(tasks, task)

    c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    status := c.DefaultQuery("status", "")

    // Simple pagination logic
    itemsPerPage := 2
    start := (page - 1) * itemsPerPage
    end := start + itemsPerPage

    var filteredTasks []models.Task
    for _, task := range tasks[start:end] {
        if status == "" || task.Status == status {
            filteredTasks = append(filteredTasks, task)
        }
    }

    c.JSON(http.StatusOK, filteredTasks)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id")
    var updatedTask models.Task
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, task := range tasks {
        if task.ID == id {
            tasks[i] = updatedTask
            c.JSON(http.StatusOK, updatedTask)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
