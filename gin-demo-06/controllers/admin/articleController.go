package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	BaseController // 继承，进行成功和失败的统一处理
}

func (article *Article) Index(c *gin.Context) {
	article.Success(c)
}

func (article *Article) Add(c *gin.Context) {
	c.String(http.StatusOK, "新增新闻-add")
}

func (article *Article) Edit(c *gin.Context) {
	c.String(http.StatusOK, "编辑新闻-edit")
}
