package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutersInit(router *gin.Engine) {
	apiRouter := router.Group("api")
	{
		apiRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "api接口")
		})
		apiRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻接口")
		})
	}
}
