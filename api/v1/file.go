package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"net/http"
	"os"
	"path"
	"server_Go/api"
	"server_Go/models"
)

func GetFileList(c *gin.Context) {
	data := models.GetFileList(c)
	c.JSON(http.StatusOK, data)
}

type File struct {
	File       string `json:"file"'`
	FileTypeId uint   `json:"file_type_id" binding:"required"`
}

func AddFile(c *gin.Context) {

	if c.Param("file_type_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"file_type_id": "缺少file_type_id",
		})
		return
	}

	SFileTypeId := c.Param("file_type_id")
	fileTypeId := cast.ToUint(SFileTypeId)

	fileType, err := models.GetFileTypeById(fileTypeId)
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	name := uuid.New().String()
	dst, err := api.UploadSingleFile(c, path.Join("file/"+fileType.Name, name))
	fmt.Println("dst", dst)
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	f := models.File{
		File:       dst,
		FileTypeId: fileTypeId,
	}

	if err := f.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, f)
}

func UpdateFile(c *gin.Context) {
	var json File
	if !api.BindAndValid(c, &json) {
		return
	}

	f, err := models.GetFileById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	name := uuid.New().String()
	dst, err := api.UploadSingleFile(c, path.Join("file", name))
	fmt.Println("dst", dst)
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	// 删除旧文件
	if json.File != "" {
		if _, err = os.Stat(json.File); err == nil {
			_ = os.Remove(json.File)
		}
	}

	n := models.File{
		File:       dst,
		FileTypeId: json.FileTypeId,
	}
	f.Updates(&n)
	c.JSON(http.StatusOK, f)
}

func DeleteFileById(c *gin.Context) {
	f, err := models.GetFileById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	// 删除旧文件
	if f.File != "" {
		if _, err = os.Stat(f.File); err == nil {
			_ = os.Remove(f.File)
		}
	}

	if err = f.Delete(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
