package routers

import (
	"GinStudy/gin-demo-04/controller/defa"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(router *gin.Engine) {
	defaultRouters := router.Group("/") // 注意花括号的位置，在下一行
	{
		defaultCon := &defa.Defa{}
		defaultRouters.GET("/", defaultCon.Index)
		defaultRouters.GET("/news", defaultCon.DefaNews)
	}
}
