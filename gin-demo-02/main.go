package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Code   int    `json:"code"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "hello world")
	})

	// 以JSON形式返回数据
	r.PUT("/json1", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"name":   "json1",
			"age":    21,
			"gender": true,
		})
	})

	// gin.H等价于map[string]interface{}类型
	r.POST("/json2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":   "json2",
			"age":    1,
			"gender": true,
		})
	})

	// 返回一个对象
	book := Book{"射雕英雄传", "金庸", 122123}
	r.DELETE("/json3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"book": book,
			"isok": true,
		})
	})

	// 可以直接写这个对象
	r.GET("/json4", func(c *gin.Context) {
		a := Book{"CSAPP", "json", 12345}
		c.JSON(200, a)
	})

	// jsonp形式的调用
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"book": book,
		})
	})

	// xml形式的调用
	r.GET("/getxml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"book":   book,
			"status": true,
		})
	})

	// html模板渲染：访问路由时加载html代码，将参数传入
	r.LoadHTMLGlob("./templates/*")
	r.GET("/gethtml", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "我是后台的数据",
		})
	})

	r.GET("/gethtml2", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "我是后台的商品数据",
		})
	})

	r.Run(":8080")

}
