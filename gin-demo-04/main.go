package main

import (
	"GinStudy/gin-demo-04/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化路由
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.Run(":8090")

}
