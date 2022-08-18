package handler

import (
	"go-todo/entity"
	uerr "go-todo/err"
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
		uerr.NotFoundErr(c)
		return
	}
	if err != nil {
		uerr.InternalErr(c, err)
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
		uerr.BadRequestErr(c, err, "Id invalid")
		return
	}

	task, err := entity.GetById(id)
	if err == gorm.ErrRecordNotFound {
		uerr.NotFoundErr(c)
		return
	}
	if err != nil {
		uerr.InternalErr(c, err)
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
		uerr.BadRequestErr(c, err, "Data invalid")
		return
	}
	err := entity.CreateTask(data)
	if err != nil {
		uerr.InternalErr(c, err)
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
		uerr.BadRequestErr(c, err, "Id invalid")
		return
	}

	err = entity.DeleteById(id)
	if err != nil {
		uerr.InternalErr(c, err)
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
		uerr.BadRequestErr(c, err, "Id invalid")
		return
	}

	err = entity.DoneTaskById(id)
	if err != nil {
		uerr.InternalErr(c, err)
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
		uerr.BadRequestErr(c, err, "Id invalid")
		return
	}

	err = entity.RejectTaskById(id)
	if err != nil {
		uerr.InternalErr(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Task has been rejected",
	})
}

func TestOrm(c *gin.Context) {
	tasks, err := entity.TestOrm()
	if len(tasks) == 0 {
		uerr.NotFoundErr(c)
		return
	}
	if err != nil {
		uerr.InternalErr(c, err)
	}
	resp := respmodel.NewListTask(tasks)
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}
