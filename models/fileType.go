package models

import "github.com/gin-gonic/gin"

type FileType struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}

func GetAllFileType() (list []FileType) {

	result := db.Model(&FileType{})

	result.Find(&list)
	return
}

func GetFileTypeList(c *gin.Context) (data *DataList) {
	var total int64
	var list []FileType

	result := db.Model(&FileType{})

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&list)

	data = GetListWithPagination(&list, c, total)

	return
}

func GetFileTypeById(id interface{}) (n FileType, err error) {
	err = db.First(&n, id).Error
	return
}

func (m *FileType) Create() error {
	return db.Create(m).Error
}

func (m *FileType) Updates(n *FileType) {
	db.Model(&FileType{}).Where("id = ?", m.ID).Updates(n)
	db.Model(&FileType{}).First(m, m.ID)
}

func (m *FileType) Delete() error {
	return db.Delete(m).Error
}
