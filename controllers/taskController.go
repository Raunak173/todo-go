package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raunak173/go-todo/initializers"
	"github.com/raunak173/go-todo/models"
)

// This is used for validation logic
var validate = validator.New()

// TaskRequestBody struct is used for incoming request bodies
type TaskRequestBody struct {
	Heading     string `json:"heading" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"is_completed"`
}

// TaskCreate creates a new task for the logged-in user
func TaskCreate(c *gin.Context) {
	var body TaskRequestBody

	// Bind JSON
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate body
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	// Get the logged-in user
	user, _ := c.Get("user")

	// Create a new task associated with the user
	task := models.Task{
		Heading:     body.Heading,
		Description: body.Description,
		IsCompleted: body.IsCompleted,
		UserID:      user.(models.User).ID, // Associate task with the logged-in user
	}

	// Save the task to the DB
	result := initializers.Db.Create(&task)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Success
	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

// GetTasks fetches all tasks for the logged-in user
func GetTasks(c *gin.Context) {
	// Get the logged-in user
	user, _ := c.Get("user")

	// Find all tasks for the user
	var tasks []models.Task
	initializers.Db.Where("user_id = ?", user.(models.User).ID).Find(&tasks)

	// Success
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

// GetTaskById fetches a specific task by its ID, only if it belongs to the logged-in user
func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")

	var task models.Task
	result := initializers.Db.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&task)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

// UpdateTask updates a task's details, only if it belongs to the logged-in user
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")

	var body TaskRequestBody

	// Bind JSON
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the task to update
	var task models.Task
	result := initializers.Db.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&task)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Update the task
	initializers.Db.Model(&task).Updates(models.Task{
		Heading:     body.Heading,
		Description: body.Description,
		IsCompleted: body.IsCompleted,
	})

	// Success
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

// MarkTaskDone marks a task as completed, only if it belongs to the logged-in user
func MarkTaskDone(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")

	// Find the task
	var task models.Task
	result := initializers.Db.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&task)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Mark the task as done
	initializers.Db.Model(&task).Update("is_completed", true)

	// Success
	c.JSON(http.StatusOK, gin.H{
		"message": "Marked the task as done",
	})
}

// DeleteTask deletes a task, only if it belongs to the logged-in user
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	user, _ := c.Get("user")

	// Find the task
	var task models.Task
	result := initializers.Db.Where("id = ? AND user_id = ?", id, user.(models.User).ID).First(&task)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Delete the task
	initializers.Db.Delete(&task)

	// Success
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted",
	})
}
