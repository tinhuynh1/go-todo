package main

import (
	"go-todo/db"
	"go-todo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()
	todo := r.Group("/api/task")
	{
		todo.GET("", handler.GetAllTask)
		todo.GET("/:id", handler.GeById)
		todo.POST("", handler.CreateTask)
		todo.DELETE("/:id", handler.GetAllTask)
		todo.PUT("/:id/done", handler.GetAllTask)
		todo.PUT("/:id/reject", handler.GetAllTask)
	}
	r.Run()
}
