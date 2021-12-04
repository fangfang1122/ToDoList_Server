package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server_Go/api"
	"server_Go/models"
	"server_Go/pkg/e"
)

func GetTeamList(c *gin.Context) {
	data := models.GetTeamList(c.Query("department_id"))
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

type Team struct {
	Name         string `json:"name" binding:"required"'`
	DepartmentId uint   `json:"department_id" binding:"required"`
}

func AddTeam(c *gin.Context) {
	var json Team
	if !api.BindAndValid(c, &json) {
		return
	}
	team := models.Team{
		Name:         json.Name,
		DepartmentId: json.DepartmentId,
	}

	if err := team.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, team)
}

func UpdateTeam(c *gin.Context) {
	var json Team
	if !api.BindAndValid(c, &json) {
		return
	}

	f, err := models.GetTeamById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.Team{
		Name:         json.Name,
		DepartmentId: json.DepartmentId,
	}
	f.Updates(&n)
	c.JSON(http.StatusOK, f)
}

func DeleteTeamById(c *gin.Context) {
	f, err := models.GetTeamById(c.Param("id"))
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
