package handler

import (
	"go-todo/entity"
	"go-todo/model"
	"go-todo/model/respmodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTask(c *gin.Context) {
	tasks, err := entity.GetAllTask()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
	}
	resp := respmodel.NewListTask(tasks)
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func GeById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	task, err := entity.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	resp := respmodel.NewTask(task)
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

func CreateTask(c *gin.Context) {
	var data model.Task
	if err := c.ShouldBind(&data); err != nil {
		panic(err)
	}
	err := entity.CreateTask(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func DeleteById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := entity.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "delete task success!",
	})
}

func DoneTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := entity.DoneTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "`",
	})
}

func RejectTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := entity.RejectTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
