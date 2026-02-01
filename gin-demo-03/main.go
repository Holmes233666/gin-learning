package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name   string
	Author string
	ISBN   string
	Year   int
}

// 待绑定的对象需要加上标签
type User struct {
	Name string `json:"username" form:"username"`
	ID   int    `json:"id" form:"id"`
}

// 待绑定的对象需要加上标签
type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

// gin框架中参数的返回、获取和绑定
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

	})

	// 加载HTML模板
	r.LoadHTMLGlob("templates/*/*.html")
	r.GET("/getdefaultindex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "这是一个html二级标题",
		})
	})

	//  get取参数：c.Query(), c.DefaultQuery()
	r.GET("/getdefaultbooks", func(c *gin.Context) {
		bookname := c.Query("bookname")
		isbn := c.DefaultQuery("isbn", "isbn-237846372")

		book := Book{bookname, "ye, her", isbn, 2001}
		c.HTML(http.StatusOK, "default/news.html", gin.H{
			"book":   book,
			"status": true,
		})
	})

	// post取参数：c.PostForm(), c.DefaultPostForm()
	r.POST("/postBook", func(c *gin.Context) {
		bookname := c.PostForm("name")
		isbn := c.DefaultPostForm("isbn", "isbn-237846372")
		book := Book{bookname, "ye, her", isbn, 2001}
		c.JSON(http.StatusOK, gin.H{
			"book":   book,
			"status": true,
		})
	})

	r.GET("/getadminindex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "这是一个html二级标题",
		})
	})

	r.GET("/getadminbooks", func(c *gin.Context) {
		book := Book{"落日高铁", "ye, her", "isbn-12uye3", 2001}
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"book":   book,
			"status": true,
		})
	})

	// 绑定参数到对象，Post和Get绑定的方式是相同的 （一般是支付场景才会使用）
	r.GET("/bindUser", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err,
			})
		}
	})

	// 从xml中解析参数
	r.POST("/bindUserFromXML", func(c *gin.Context) {
		xmlSlice, _ := c.GetRawData() // 返回byte类型切片

		a := Article{}

		if err := xml.Unmarshal(xmlSlice, &a); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"data": a,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
		}
	})

	// 动态路由传值
	// list/123 list/456 对应不同的资源
	r.GET("/dynamicRouter/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "%v", cid)
	})

	// 启动路由
	r.Run(":8080")
}
