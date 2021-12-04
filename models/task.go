package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	CreateUserId uint `json:"create_user_id" gorm:"not null,index"`
}

func GetTaskList(c *gin.Context) (data *DataList) {
	var list []Task
	result := db.Model(&Task{})


	var total int64

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&list)

	data = GetListWithPagination(&list, c, total)
	return
}

func GetTaskById(id interface{}) (n Auth, err error) {
	err = db.First(&n, id).Error
	return
}

func (f *Task) Create() error {
	return db.Create(f).Error
}

func (f *Task) Update(n *Task) {
	db.Model(&Task{}).Where("id = ?", f.ID).Updates(n)
	db.Model(&Task{}).First(f, f.ID)
}

func (f *Task) Delete() error {
	return db.Delete(f).Error
}
