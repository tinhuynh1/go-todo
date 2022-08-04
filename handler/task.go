package handler

import (
	"go-todo/entity"
	"go-todo/model"
	"go-todo/model/respmodel"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTask(c *gin.Context) {
	tasks, err := entity.GetAllTask()
	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "tasks is empty",
		})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	task, err := entity.GetById(id)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "task not found",
			"id":  id,
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
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
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}
	err := entity.CreateTask(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create task successfully!",
	})
}

func DeleteById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = entity.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been deleted",
	})
}

func DoneTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = entity.DoneTaskById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been done",
	})
}

func RejectTaskById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = entity.RejectTaskById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been rejected",
	})
}
