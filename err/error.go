package err

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalErr(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func NotFoundErr(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func BadRequestErr(c *gin.Context, err error, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
		"msg":   msg,
	})
}
