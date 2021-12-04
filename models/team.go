package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	Name         string      `json:"name"`
	CreateUserId uint `json:"create_user_id" gorm:"not null,index"`
	Users []User `json:"users"`
}

func GetAllTeam(c *gin.Context) (list []Team) {

	result := db.Model(&Team{}).Preload("Users")

	result.Find(&list)
	return
}

func GetTeamById(id interface{}) (n Team, err error) {
	err = db.First(&n, id).Error
	return
}

func (f *Team) Create() error {
	return db.Create(f).Error
}

func (f *Team) Update(n *Team) {
	db.Model(&Team{}).Where("id = ?", f.ID).Updates(n)
	db.Model(&Team{}).First(f, f.ID)
}

func (f *Team) Delete() error {
	return db.Delete(f).Error
}
