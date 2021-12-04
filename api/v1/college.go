package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_Go/api"
	"server_Go/models"
	"server_Go/pkg/e"
)

func GetCollegeList(c *gin.Context) {
	data := models.GetCollegeList()
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type College struct {
	Name string `json:"name" binding:"required"'`
}

func AddCollege(c *gin.Context) {
	var json College
	if !api.BindAndValid(c, &json) {
		return
	}
	college := models.College{
		Name: json.Name,
	}

	if err := college.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, college)
}

func UpdateCollege(c *gin.Context) {
	var json College
	if !api.BindAndValid(c, &json) {
		return
	}

	college, err := models.GetCollegeById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.College{
		Name: json.Name,
	}
	college.Updates(&n)
	c.JSON(http.StatusOK, college)
}

func DeleteCollegeById(c *gin.Context) {
	college, err := models.GetCollegeById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	if err = college.Delete(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
