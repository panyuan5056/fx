package xstrategy

import (
	"fx/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	var strategy Strategy
	code := e.INVALID_PARAMS
	if err := c.BindJSON(&strategy); err == nil {
		code = e.SUCCESS
		status, _ := strategy.valid()
		if status == e.SUCCESS {
			strategy.save()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})
	c.Abort()
}

func List(c *gin.Context) {
	result := Find()
	c.JSON(http.StatusOK, gin.H{
		"status": e.SUCCESS,
		"msg":    e.GetMsg(e.SUCCESS),
		"data":   result,
	})
	c.Abort()
}
