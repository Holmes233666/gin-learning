package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutersInit(router *gin.Engine) {
	adminRouter := router.Group("admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "admin")
		})

		adminRouter.GET("/users", func(c *gin.Context) {
			c.String(http.StatusOK, "管理员列表")
		})
		adminRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "管理员新闻主页")
		})
	}
}
