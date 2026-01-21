package main

import (
	"github.com/gin-gonic/gin"
	"net/http" // 状态码， http.Status
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 配置路由：表示访问 "/" 时触发后面的回调函数
	r.GET("/", func(c *gin.Context) {
		// String writes the given string into the response body.
		c.String(http.StatusOK, "Hello World") // 状态码
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "Hello News")
	})

	r.POST("/postidea", func(c *gin.Context) {
		c.String(200, "post请求，提交表单")
	})

	r.PUT("/putidea", func(c *gin.Context) {
		c.String(200, "put请求，新增资源")
	})

	r.DELETE("/deleteidea", func(c *gin.Context) {
		c.String(200, "delete 请求，删除资源")
	})

	// 在启动一个web服务
	r.Run(":8080") // 如果不写8080，那么默认的端口也会是8080
}
