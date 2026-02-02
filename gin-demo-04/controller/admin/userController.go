package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserIndex(c *gin.Context) {
	c.String(http.StatusOK, "admin")
}

func UserAdd(c *gin.Context) {
	c.String(http.StatusOK, "admin user add")
}

func UserEdit(c *gin.Context) {
	c.String(http.StatusOK, "admin user edit")
}
