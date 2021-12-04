package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Sponsor struct {
	gorm.Model
	Name string `json:"name"`
	Type uint   `json:"type"`
	Logo string `json:"logo"`
}

func GetSponsorById(id interface{}) (n Sponsor, err error) {
	err = db.First(&n, id).Error
	return
}

// 有分页
func GetSponsorList(c *gin.Context) (data *DataList) {
	var sponsors []Sponsor
	result := db.Model(&Sponsor{})

	if c.Query("type") != "" {
		result = result.Where("type = ?", c.Query("type"))
	}

	var total int64

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&sponsors)

	data = GetListWithPagination(&sponsors, c, total)

	result.Find(&sponsors)
	return
}

//无分页，全部按类加载
func GetAllSponsor(typeId interface{}) (list []Sponsor) {

	result := db.Model(&Sponsor{})

	if typeId != "" {
		result = result.Where("type = ?", typeId)
	}

	result.Find(&list)
	return
}

func (sponsor *Sponsor) Create() error {
	return db.Create(sponsor).Error
}

func (sponsor *Sponsor) Updates(n *Sponsor) {
	db.Model(&Sponsor{}).Where("id = ?", sponsor.ID).Updates(n)
	db.Model(&Sponsor{}).First(sponsor, sponsor.ID)
}

func (m *Sponsor) Delete() error {
	return db.Delete(m).Error
}
