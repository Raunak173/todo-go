package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raunak173/go-todo/controllers"
	"github.com/raunak173/go-todo/initializers"
	"github.com/raunak173/go-todo/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {

	r := gin.Default()
	//Task controllers
	r.POST("/task/create", middleware.RequireAuth, controllers.TaskCreate)
	r.GET("/tasks", middleware.RequireAuth, controllers.GetTasks)
	r.GET("/task/:id", middleware.RequireAuth, controllers.GetTaskById)
	r.PUT("/task/:id", middleware.RequireAuth, controllers.UpdateTask)
	r.POST("/task/:id", middleware.RequireAuth, controllers.MarkTaskDone)
	r.DELETE("/task/:id", middleware.RequireAuth, controllers.DeleteTask)
	//User controllers
	r.POST("/user/signup", controllers.SignUp)
	r.POST("/user/login", controllers.Login)
	r.Run()
}
