package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	fmt.Println("中间件1")
	c.Set("userId", "123")
	c.Next()
	// 定义一个goroutine统计日志
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path" + cCp.Request.URL.Path)
	}()
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
