package routers

import (
	"GinStudy/gin-demo-04/controller/admin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutersInit(router *gin.Engine) {
	adminRouter := router.Group("admin")
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "admin")
		})

		// 非结构体的写法
		adminRouter.GET("/user", admin.UserIndex) // 不写括号，代表注册函数，而非调用，真正调用是Gin框架调用的
		adminRouter.GET("/user/add", admin.UserAdd)
		adminRouter.GET("/user/edit", admin.UserEdit)

		article := &admin.Article{}
		// 结构体的写法
		adminRouter.GET("/article", article.Index)
		adminRouter.GET("/article/add", article.Add)
		adminRouter.GET("/article/edit", article.Edit)
	}
}
