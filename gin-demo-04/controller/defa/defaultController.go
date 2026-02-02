package defa

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Defa struct {
}

func (api *Defa) Index(c *gin.Context) {
	c.String(http.StatusOK, "api接口")
}

func (api *Defa) DefaNews(c *gin.Context) {
	c.String(http.StatusOK, "api新闻接口")
}
