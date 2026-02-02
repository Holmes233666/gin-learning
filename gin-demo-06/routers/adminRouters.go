package routers

import (
	"GinStudy/gin-demo-06/controllers/admin"
	"GinStudy/gin-demo-06/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutersInit(router *gin.Engine) {
	// 分组路由中配置中间件的方式1:
	//adminRouter := router.Group("admin", middlewares.InitMiddleWare)
	adminRouter := router.Group("admin")
	adminRouter.Use(middlewares.InitMiddleWare) // 分组路由中配置中间件的第二种方式
	{
		adminRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "admin")
		})

		// 非结构体的写法
		adminRouter.GET("/user", admin.UserIndex) // 不写括号，代表注册函数，而非调用，真正调用是Gin框架调用的
		adminRouter.GET("/user/add", admin.UserAdd)
		adminRouter.GET("/user/edit", admin.UserEdit)
		adminRouter.POST("/user/doBulkUpload", admin.DoBulkUpload) // 上传多个同名文件
		adminRouter.POST("/user/doUpload", admin.DoUpload)         // 上传文件

		article := &admin.Article{}
		// 结构体的写法
		adminRouter.GET("/article", article.Index)
		adminRouter.GET("/article/add", article.Add)
		adminRouter.GET("/article/edit", article.Edit)
	}
}
