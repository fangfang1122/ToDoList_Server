package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

type Auth struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null"`
	Username   string `json:"username" gorm:"unique"`
	Password   string `json:"-" gorm:"not null"`
	CreateById int    `json:"create_by_id" `
}

func GetAuthById(id interface{}) (n Auth, err error) {
	err = db.First(&n, id).Error
	return
}

func FindAuth(username string) (auth Auth, err error) {
	result := db.Where(&Auth{Username: username}).First(&auth)
	err = result.Error
	if err != nil {
		log.Println("用户不存在")
	}
	return
}

func GetAuthList(c *gin.Context) (data *DataList) {
	var total int64
	var auths []Auth

	result := db.Model(&Auth{})

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&auths)

	data = GetListWithPagination(&auths, c, total)

	return
}

func ExistAuthByName(username string) bool {
	var auth Auth
	db.Where("username = ?", username).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func (auth *Auth) Save() error {
	return db.Create(auth).Error
}

func (auth *Auth) Updates(n *Auth) {
	db.Model(&Auth{}).Where("id = ?", auth.ID).Updates(n)
	db.Model(&Auth{}).First(auth, auth.ID)
}

func (auth *Auth) Delete() error {
	return db.Delete(auth).Error
}
