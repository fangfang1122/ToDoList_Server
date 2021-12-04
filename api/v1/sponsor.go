package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path"
	api "server_Go/api"
	"server_Go/models"
	"server_Go/pkg/e"
)

type Sponsor struct {
	Name string `json:"name" binding:"required"`
	Type uint   `form:"type" json:"type" binding:"required"`
	Logo string `json:"logo"`
}

func AddSponsor(c *gin.Context) {
	var json Sponsor
	if !api.BindAndValid(c, &json) {
		return
	}

	sponsor := models.Sponsor{
		Name: json.Name,
		Type: json.Type,
	}
	err := sponsor.Create()
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": sponsor,
	})
}

func GetSponsorList(c *gin.Context) {
	data := models.GetSponsorList(c)
	c.JSON(http.StatusOK, data)
}

func GetAllSponsor(c *gin.Context) {
	data := models.GetAllSponsor(c.Query("type"))
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func EditSponsor(c *gin.Context) {
	var json Sponsor
	if !api.BindAndValid(c, &json) {
		return
	}
	sponsor, err := models.GetSponsorById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.Sponsor{
		Name: json.Name,
		Type: json.Type,
		Logo: json.Logo,
	}
	sponsor.Updates(&n)
	c.JSON(http.StatusOK, sponsor)
}

func DeleteSponsorById(c *gin.Context) {

	sponsor, err := models.GetSponsorById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	// 删除旧文件
	if sponsor.Logo != "" {
		if _, err = os.Stat(sponsor.Logo); err == nil {
			_ = os.Remove(sponsor.Logo)
		}
	}

	if err = sponsor.Delete(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func UploadSponsorLogo(c *gin.Context) {

	sponsor, err := models.GetSponsorById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	name := uuid.New().String()
	dst, err := api.UploadSingleFile(c, path.Join("logo", name))
	fmt.Println("dst", dst)
	if err != nil {
		api.ErrHandler(c, err)
		return
	}

	// 删除旧照片
	if sponsor.Logo != "" {
		if _, err = os.Stat(sponsor.Logo); err == nil {
			_ = os.Remove(sponsor.Logo)
		}
	}

	n := models.Sponsor{
		Name: sponsor.Name,
		Type: sponsor.Type,
		Logo: dst,
	}
	sponsor.Updates(&n)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     dst,
		"data":    n,
	})
}
