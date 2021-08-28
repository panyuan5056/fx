package core

import (
	"fx/core/xdb"
	"fx/core/xfile"
	"fx/core/xstrategy"

	"github.com/gin-gonic/gin"
)

func InitApiv1(apiv1 *gin.RouterGroup) *gin.RouterGroup {
	{
		apiv1.POST("/get", Get)
		apiv1.POST("/db/find", xdb.Find)
		apiv1.POST("/file/find", xfile.Find)
		apiv1.POST("/strategy/add", xstrategy.Add)
		apiv1.POST("/strategy/list", xstrategy.List)
	}
	return apiv1
}
