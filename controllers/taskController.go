package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/raunak173/go-todo/initializers"
	"github.com/raunak173/go-todo/models"
)

// This is used for our validation logic
var validate = validator.New()

// A struct of taskBody that is use to be given by the user as body parameter
type TaskRequestBody struct {
	Heading     string `json:"heading" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsCompleted bool   `json:"is_completed"`
}

func TaskCreate(c *gin.Context) {

	var body TaskRequestBody

	//Binding json
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//Validating body
	if err := validate.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	//Getting data from req body
	task := models.Task{
		Heading:     body.Heading,
		Description: body.Description,
		IsCompleted: body.IsCompleted,
	}

	//Creating a task and savin it in DB
	result := initializers.Db.Create(&task)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//Success
	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func GetTasks(c *gin.Context) {

	//Find all
	var tasks []models.Task
	initializers.Db.Find(&tasks)

	//Success
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func GetTaskById(c *gin.Context) {

	id := c.Param("id")

	var task models.Task
	initializers.Db.First(&task, id)

	//Success
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Heading     string `json:"heading"`
		Description string `json:"description"`
		IsCompleted bool
	}

	//Binding json
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	//Finding the task by id we want to update
	var task models.Task
	initializers.Db.First(&task, id)

	initializers.Db.Model(&task).Updates(models.Task{
		Heading:     body.Heading,
		Description: body.Description,
		IsCompleted: body.IsCompleted,
	})
	//Success
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func MarkTaskDone(c *gin.Context) {

	id := c.Param("id")

	var task models.Task
	initializers.Db.First(&task, id)

	initializers.Db.Model(&task).Updates(models.Task{
		IsCompleted: true,
	})

	//Success
	c.JSON(http.StatusOK, gin.H{
		"message": "Marked the task as done",
	})
}

func DeleteTask(c *gin.Context) {

	id := c.Param("id")

	initializers.Db.Delete(&models.Task{}, id)

	//Success
	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted",
	})
}
