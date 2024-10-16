package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raunak173/go-todo/controllers"
	"github.com/raunak173/go-todo/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {

	r := gin.Default()
	r.POST("/task/create", controllers.TaskCreate)
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/task/:id", controllers.GetTaskById)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.POST("/task/:id", controllers.MarkTaskDone)
	r.DELETE("/task/:id", controllers.DeleteTask)
	r.Run()
}
