package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Type struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	CreateUserId uint `json:"create_user_id" gorm:"not null,index"`
}

func GetAllType(c *gin.Context) (list []Type) {

	result := db.Model(&Type{})

	result.Find(&list)
	return
}

func GetTypeById(id interface{}) (n Auth, err error) {
	err = db.First(&n, id).Error
	return
}

func (f *Type) Create() error {
	return db.Create(f).Error
}

func (f *Type) Update(n *Type) {
	db.Model(&Type{}).Where("id = ?", f.ID).Updates(n)
	db.Model(&Type{}).First(f, f.ID)
}

func (f *Type) Delete() error {
	return db.Delete(f).Error
}
