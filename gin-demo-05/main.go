package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	fmt.Println("中间件1")
	c.Next()
	fmt.Println("中间件1结束")
}

func InitMiddleWare2(c *gin.Context) {
	fmt.Println("中间件2")
	c.Next()
	fmt.Println("中间件2结束")
}

func InitMiddleWare3(c *gin.Context) {
	fmt.Println("中间件3")
	c.Next()
	fmt.Println("中间件3结束")
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	r.GET("/admin",
		InitMiddleWare, InitMiddleWare2, InitMiddleWare3,
		func(c *gin.Context) {
			time.Sleep(1 * time.Second)
			c.String(http.StatusOK, "admin")
		})

	r.Run(":8080")
}
