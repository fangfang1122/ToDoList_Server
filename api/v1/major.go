package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_Go/api"
	"server_Go/models"
	"server_Go/pkg/e"
)

func GetMajorList(c *gin.Context) {
	data := models.GetMajorList(c.Query("college_id"))
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type Major struct {
	Name      string `json:"name" binding:"required"'`
	CollegeId uint   `json:"college_id" binding:"required"`
}

func AddMajor(c *gin.Context) {
	var json Major
	if !api.BindAndValid(c, &json) {
		return
	}

	major := models.Major{
		Name:      json.Name,
		CollegeId: json.CollegeId,
	}

	if err := major.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, major)

}

func UpdateMajor(c *gin.Context) {
	var json Major
	if !api.BindAndValid(c, &json) {
		return
	}

	major, err := models.GetMajorById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.Major{
		Name:      json.Name,
		CollegeId: json.CollegeId,
	}
	major.Updates(&n)
	c.JSON(http.StatusOK, major)
}

func DeleteMajorById(c *gin.Context) {
	major, err := models.GetMajorById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	if err = major.Delete(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
