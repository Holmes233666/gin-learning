package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(InitMiddleWare, InitMiddleWare2, InitMiddleWare3) // r.Use()可以使用全局中间件
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "admin")
	})

	r.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "admin")
	})

	r.Run(":8080")
}
