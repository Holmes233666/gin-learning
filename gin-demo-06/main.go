package main

import (
	"GinStudy/gin-demo-06/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("gin-demo-06/templates/admin/*.html")
	r.Static("/static", "./gin-demo-06/static")
	// 初始化路由
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.Run(":8090")

}
