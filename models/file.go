package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type File struct {
	gorm.Model
	File       string    `json:"file" gorm:"not null"`
	FileTypeId uint      `json:"file_type_id" gorm:"not null"`
	FileType   *FileType `json:"file_type,omitempty" `
}

func GetAllFile(typeId interface{}) (list []File) {

	result := db.Model(&File{}).Preload("FileType")

	if typeId != "" {
		result = result.Where("type_id = ?", typeId)
	}

	result.Find(&list)
	return
}

func GetFileList(c *gin.Context) (data *DataList) {
	var total int64
	var list []File

	result := db.Model(&File{}).Preload("FileType")

	if c.Query("file_type_id") != "" {
		result = result.Where("file_type_id = ?", c.Query("file_type_id"))
	}

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&list)

	data = GetListWithPagination(&list, c, total)

	return
}

func GetFileById(id interface{}) (n File, err error) {
	err = db.First(&n, id).Error
	return
}

func (m *File) Create() error {
	return db.Create(m).Error
}

func (m *File) Updates(n *File) {
	db.Model(&File{}).Where("id = ?", m.ID).Updates(n)
	db.Model(&File{}).First(m, m.ID)
}

func (m *File) Delete() error {
	return db.Delete(m).Error
}
