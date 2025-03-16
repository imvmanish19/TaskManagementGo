package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"task-service/internal/models"
	"task-service/internal/httpclient"
	"github.com/gin-gonic/gin"
)

// Global in-memory tasks and userServiceClient
var tasks []models.Task
var userServiceClient = httpclient.NewUserServiceClient()

// CreateTask handles creating a new task
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user information from User Service
	user, err := userServiceClient.GetUser(task.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	fmt.Printf("Task: %v", task)
	fmt.Printf("User: %v", user)

	// Assign an ID and append to tasks
	task.ID = strconv.Itoa(len(tasks) + 1)
	task.Status = "Pending" // Default status can be "Pending"

	// Store the task (this is just an in-memory store for now)
	tasks = append(tasks, task)

	// Return task in the response (without nested user details)
	c.JSON(http.StatusCreated, task)
}

// GetTasks retrieves a list of tasks with pagination and optional filtering by status
func GetTasks(c *gin.Context) {
    userID, _ := strconv.Atoi(c.DefaultQuery("user_id", "0"))

    if userID < 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "User id is not valid"})
        return
    }
    
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    status := c.DefaultQuery("status", "")

    // Pagination settings
    itemsPerPage := 2
    start := (page - 1) * itemsPerPage
    end := start + itemsPerPage

    // Ensure end index does not exceed slice bounds
    if end > len(tasks) {
        end = len(tasks)
    }

    // Filter tasks based on the status and user_id query parameters
    var filteredTasks []models.Task
    for _, task := range tasks[start:end] {
        if (status == "" || task.Status == status) && (userID == 0 || task.UserID == strconv.Itoa(userID)) {
            filteredTasks = append(filteredTasks, task)
        }
    }

    // Prepare the response with filtered tasks
    c.JSON(http.StatusOK, filteredTasks)
}

// DeleteTask handles the deletion of a task by ID
func DeleteTask(c *gin.Context) {
    taskID := c.Param("id")
    userID := c.DefaultQuery("user_id", "")  // Assuming user_id is passed as a query parameter or via JWT in headers

    // Ensure that userID is not empty and matches the task's UserID
    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
        return
    }

    // Find the index of the task to delete
    var taskIndex int
    var found bool
    for i, task := range tasks {
        if task.ID == taskID {
            if task.UserID != userID {
                // If the user ID doesn't match the task's UserID, deny the action
                c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this task"})
                return
            }
            taskIndex = i
            found = true
            break
        }
    }

    // If task is not found, return an error
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Task with ID %s not found", taskID)})
        return
    }

    // Delete the task by removing it from the slice
    tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Task with ID %s has been deleted", taskID),
    })
}

// UpdateTask handles updating a task by ID
func UpdateTask(c *gin.Context) {
    taskID := c.Param("id")
    userID := c.DefaultQuery("user_id", "")  // Assuming user_id is passed as a query parameter or via JWT in headers
    var updatedTask models.Task

    // Ensure that userID is not empty and matches the task's UserID
    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
        return
    }

    // Bind the JSON body to the task struct
    if err := c.ShouldBindJSON(&updatedTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find the task in the list
    var taskIndex int
    var found bool
    for i, task := range tasks {
        if task.ID == taskID {
            if task.UserID != userID {
                // If the user ID doesn't match the task's UserID, deny the action
                c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this task"})
                return
            }
            taskIndex = i
            found = true
            break
        }
    }

    // If task is not found, return an error
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Task with ID %s not found", taskID)})
        return
    }

    // Update the task
	updatedTask.ID = taskID
    tasks[taskIndex] = updatedTask

    c.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("Task with ID %s has been updated", taskID),
        "task":    updatedTask,
    })
}
