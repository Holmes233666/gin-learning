package routers

import (
	"GinStudy/gin-demo-06/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		a := &api.Api{}
		apiRouter.GET("/", a.Index)
		apiRouter.GET("/news", a.ApiNews)
	}
}
