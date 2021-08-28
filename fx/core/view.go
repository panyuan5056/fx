package core

import (
	"fmt"
	"fx/models"
	"fx/pkg/e"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	var get GetParm
	code := e.INVALID_PARAMS
	var queue models.Queue
	if err := c.BindJSON(&get); err == nil {
		code = e.SUCCESS
		queue = models.Get(get.Id)
	} else {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": queue,
	})
	c.Abort()
}
