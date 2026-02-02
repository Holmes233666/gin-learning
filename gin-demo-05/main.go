package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	r.GET("/admin",
		func(c *gin.Context) { fmt.Printf("aaa") },
		func(c *gin.Context) { c.String(http.StatusOK, "admin") },
	)
	
	r.Run(":8080")
}
