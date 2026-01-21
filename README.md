# gin-learning
# Gin + Gorm + Go RPC学习

## 一、Gin框架

Gin：高性能的HTTP Web框架，Gin 专为构建 REST API、Web 应用程序和微服务而设计，其中速度和开发人员生产力至关重要。

官网：https://gin-gonic.com

Gin的安装：项目目录下使用`go get -u github.com/gin-gonic/gin`进行下载和安装

### demo

```go
func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 配置路由：表示访问 "/" 时触发后面的回调函数
	r.GET("/", func(c *gin.Context) {
		// String writes the given string into the response body.
		c.String(200, "Hello World")
	})

	// 启动一个web服务
	r.Run(":8080") // 如果不写8080，那么默认的端口也会是8080
}

```

在本地访问：localhost:8080/

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20260120154124752.png" alt="image-20260120154124752" style="zoom:50%;" />

路由可配置多个：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"	// 状态码， http.Status
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 配置路由：表示访问 "/" 时触发后面的回调函数
	r.GET("/", func(c *gin.Context) {
		// String writes the given string into the response body.
		c.String(http.StatusOK, "Hello World")	// 状态码
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
```

#### Go语言程序的热加载

gin框架中并没有实现热加载的功能：即修改代码后需要重启才能使用新的服务。但是beego框架中提供了热加载的功能：

`go install github.com/pilu/fresh`安装fresh包，使用`fresh`命令即可启动热加载。之后对代码的任意改变均不用重启项目即可实现。

> `go get`和`go install`在`go 1.17`更新后功能分离，`go get` 现在只负责处理代码依赖（库），而 `go install` 负责安装可执行工具（软件）。所以使用外部程序必须使用go install配置命令的环境变量到$GOPATH/bin下，否则可执行工具无法使用。

![image-20260120170842838](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20260120170842838.png)

#### gin.context的返回类型

（1）`c.String(), c.JSON(), c.JSONP()`

返回类型可以是：`c.String(), c.JSON(), c.JSONP()`，具体使用如下：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name   string
	Author string
	Code   int
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
		a := &Book{"CSAPP", "json", 12345}
		c.JSON(200, a)
	})

	r.Run(":8080")

}
```

其中数据类型`gin.H`表示的就是`map[string]interface{}`。另外，3和4返回的结果有所不同：

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20260120181933602.png" alt="image-20260120181933602" style="zoom:25%;" />

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20260120182215614.png" alt="image-20260120182215614" style="zoom:25%;" />

`c.JSON` 的核心作用就是封装了 Go 标准库 `encoding/json` 的序列化逻辑，并自动处理 HTTP 响应头。

当调用 `c.JSON(code, obj)` 时，Gin 内部实际上做了这三件事：

- **设置响应头**：自动将 `Content-Type` 设置为 `application/json; charset=utf-8`。
- **状态码**：将 HTTP 状态码写入响应（即传入的 `http.StatusOK` 或 `200`）。
- **序列化 (Marshal)**：调用标准库 `encoding/json` 的逻辑，把传入的 `obj`（无论是 `map`、`gin.H` 还是 `struct`）转换成 JSON 格式的字节流，并写入到 `http.ResponseWriter` 中。

两个必须注意的细节：它底层是 `json.Marshal`，那么就要遵守 Go 标准库的序列化规则：

**字段必须导出（首字母大写）**： 如果 `Book` 结构体字段是 `name` (小写)，即使加了 tag，`c.JSON` 也会输出空值或者忽略它。因为 `json` 包通过反射（Reflection）访问字段，无法访问私有字段。

- `name string` -> JSON 中看不到
- `Name string` -> JSON 中看得到

  `gin.H` 是 `map[string]interface{}` 的别名（源码里写着 `type H map[string]any`）。Gin 发明这个缩写只是为了让在写嵌套 JSON 时少打几个字，它的序列化过程和普通的 map 完全一样。

----

`jsonp`与与`json`的区别是可以传入回调函数，解决跨域问题。

![image-20260120185138579](https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20260120185138579.png)

（2）返回形式是`html`和`xml`：`c.HTML(), c.XML()`

使用方法如下：

```go
// xml形式的调用
	r.GET("/getxml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"book":   book,
			"status": true,
		})
	})

	// html模板渲染：访问路由时加载html代码，将参数传入
	r.LoadHTMLGlob("./templates/*")	// 注意使用前需要加载Html模板
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
```

<img src="https://cdn.jsdelivr.net/gh/Holmes233666/blogImage/images/image-20260121151809986.png" alt="image-20260121151809986" style="zoom:25%;" />

