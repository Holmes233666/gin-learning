package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRoutersInit(router *gin.Engine) {
	defaultRouters := router.Group("/") // 注意花括号的位置，在下一行
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻首页")
		})
	}
}
