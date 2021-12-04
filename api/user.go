package api

import (
	"ToDoList_Go/models"
	"ToDoList_Go/pkg/e"
	"ToDoList_Go/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username   string `json:"username" binding:"required"`
	Avatar string `json:"avatar"`
}

func Login(c *gin.Context) {

	var json struct{
		Code string `json:"code" binding:"required,min=6"`
	}

	if !BindAndValid(c, &json) {
		return
	}

	ok,openid:=GetOpenId(c,json.Code)
	if !ok{
		return
	}
	fmt.Println(openid)
	user, err := models.GetUserByOpenId(openid)
	if err != nil {
		ErrHandler(c, err)
		return
	}
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
	} else {
		data["token"] = token
		code = e.SUCCESS
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

func UserInfo(c *gin.Context) {
	user := CurrentUser(c)
	c.JSON(http.StatusOK, user)
}


func EditAuth(c *gin.Context) {
	var json User
	if !BindAndValid(c, &json) {
		return
	}

	f, err := models.GetUserById(c.Param("id"))
	if err != nil {
		ErrHandler(c, err)
		return
	}

	n := models.User{
		Username: json.Username,
		Avatar: json.Avatar,
	}
	f.Update(&n)
	c.JSON(http.StatusOK, f)
}



