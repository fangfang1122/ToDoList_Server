package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_Go/api"
	"server_Go/models"
	"server_Go/pkg/e"
)

func GetDepartmentList(c *gin.Context) {
	data := models.GetDepartmentList()
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type Department struct {
	Name string `json:"name" binding:"required"'`
}

func AddDepartment(c *gin.Context) {
	var json Department
	if !api.BindAndValid(c, &json) {
		return
	}
	f := models.Department{
		Name: json.Name,
	}

	if err := f.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, f)
}

func UpdateDepartment(c *gin.Context) {
	var json Department
	if !api.BindAndValid(c, &json) {
		return
	}

	f, err := models.GetDepartmentById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.Department{
		Name: json.Name,
	}
	f.Updates(&n)
	c.JSON(http.StatusOK, f)
}

func DeleteDepartmentById(c *gin.Context) {
	f, err := models.GetDepartmentById(c.Param("id"))
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
