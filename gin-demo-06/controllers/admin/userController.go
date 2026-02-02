package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func UserIndex(c *gin.Context) {
	if userId, exist := c.Get("userId"); exist {
		id, err := userId.(string)
		if !err {
			c.String(http.StatusOK, "请求失败")
		} else {
			c.String(http.StatusBadRequest, id)
		}
	} else {
		c.String(http.StatusBadRequest, "请求失败")
	}
}

func UserAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

func UserEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}

func DoUpload(c *gin.Context) {
	userName := c.PostForm("username")
	file, err := c.FormFile("face")
	if err != nil {
		c.String(http.StatusBadRequest, "获取上传文件失败：%s", err.Error())
		return
	}
	dst := path.Join("./gin-demo-06/static/upload", file.Filename) // 当前目录其实是引入到main.go之后的目录
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.String(http.StatusBadRequest, "保存文件失败：%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userName": userName,
		"dst":      dst,
	})
}

func DoBulkUpload(c *gin.Context) {
	userName := c.PostForm("username")
	// 读取多个相同名字的文件
	form, _ := c.MultipartForm()
	files, _ := form.File["face[]"] // files：[]*multipart.FileHeader类型

	for _, file := range files {
		dst := path.Join("./gin-demo-06/static/upload", file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			c.String(http.StatusBadRequest, "文件报错失败：%s", err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"userName": userName,
		"success":  true,
	})
}
