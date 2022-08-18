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
		todo.DELETE("/:id", handler.DeleteById)
		todo.PUT("/:id/done", handler.DoneTaskById)
		todo.PUT("/:id/reject", handler.RejectTaskById)
		todo.GET("/test", handler.TestOrm)
	}
	r.Run()
}
