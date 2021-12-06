package v1

import (
	"ToDoList_Go/api"
	"ToDoList_Go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type List struct {
	Name string `json:"name" binding:"required"'`
}

func GetAllList(c *gin.Context) {
	user := api.CurrentUser(c)
	data := models.GetAllList(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func AddList(c *gin.Context) {
	var json List
	if !api.BindAndValid(c, &json) {
		return
	}
	user := api.CurrentUser(c)
	if user.CreateListAmount >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "每位用户创建清单数量不能超过10个",
		})
		return
	}

	f := models.List{
		Name:   json.Name,
		UserId: user.ID,
	}

	if err := f.Create(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	n := models.User{
		CreateListAmount: user.CreateListAmount + 1,
	}
	user.Update(&n)
	c.JSON(http.StatusOK, f)

}

func UpdateList(c *gin.Context) {
	var json List
	if !api.BindAndValid(c, &json) {
		return
	}

	f, err := models.GetListById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	user := api.CurrentUser(c)
	if f.UserId != user.ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无法修改非自己创建的清单",
		})
		return
	}
	n := models.List{
		Name: json.Name,
	}
	f.Update(&n)
	c.JSON(http.StatusOK, f)
}

func DeleteList(c *gin.Context) {
	f, err := models.GetListById(c.Param("id"))
	if err != nil {
		api.ErrHandler(c, err)
		return
	}
	user := api.CurrentUser(c)
	if f.UserId != user.ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "无法删除非自己创建的清单",
		})
		return
	}

	n := models.User{
		CreateListAmount: user.CreateListAmount - 1,
	}
	user.Update(&n)

	if err = f.Delete(); err != nil {
		api.ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
