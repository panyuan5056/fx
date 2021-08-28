package xdb

import (
	"fx/models"
	"fx/pkg/e"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) {
	var config Config
	code := e.INVALID_PARAMS
	data := map[string]int{}
	if c.BindJSON(&config) == nil {
		//将数据加入到异步队列
		code = e.SUCCESS
		if content, ok := config.str(); ok {
			id := models.Push("1", content)
			data["id"] = id
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": data,
	})

	c.Abort()
}
