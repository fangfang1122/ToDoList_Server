package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_Go/api"
	"server_Go/models"
)

func GetFileTypeList(c *gin.Context) {
	data := models.GetFileTypeList(c)
	c.JSON(http.StatusOK, data)
}

type FileType struct {
	Name string `json:"name" binding:"required"'`
}

func AddFileType(c *gin.Context) {
	var json FileType
	if !api.BindAndValid(c, &json) {
		return
	}
	f := models.FileType{
		Name: json.Name,
	}

	if err := f.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, f)
}

func UpdateFileType(c *gin.Context) {
	var json FileType
	if !api.BindAndValid(c, &json) {
		return
	}

	f, err := models.GetFileTypeById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.FileType{
		Name: json.Name,
	}
	f.Updates(&n)
	c.JSON(http.StatusOK, f)
}

func DeleteFileTypeById(c *gin.Context) {
	f, err := models.GetFileTypeById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	if err = f.Delete(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
