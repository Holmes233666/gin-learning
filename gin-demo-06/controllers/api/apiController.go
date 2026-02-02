package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
}

func (api *Api) Index(c *gin.Context) {
	c.String(http.StatusOK, "api接口")
}

func (api *Api) ApiNews(c *gin.Context) {
	c.String(http.StatusOK, "api新闻接口")
}
