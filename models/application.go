package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Application struct {
	gorm.Model

	Name     string `json:"name"`
	Gender   uint   `json:"gender"`
	SchoolId string `json:"school_id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Ability  string `json:"ability"`
	Knowing  string `json:"knowing"`
	Why      string `json:"why"`

	CollegeId    int         `json:"college_id"`
	College      *College    `json:"college,omitempty"`
	MajorId      int         `json:"major_id"`
	Major        *Major      `json:"major,omitempty"`
	DepartmentId int         `json:"department_id"`
	Department   *Department `json:"department,omitempty"`
	TeamId       int         `json:"team_id"`
	Team         *Team       `json:"team,omitempty"`
}

func GetApplicationList(c *gin.Context, CollegeId, MajorId, DepartmentId, TeamId interface{}) (data *DataList) {

	var total int64
	var applications []Application

	result := db.Model(&Application{}).Preload("Department").Preload("College").Preload("Major").Preload("Team")

	if CollegeId != "" {
		result = result.Where(" college_id = ?", CollegeId)
	}

	if MajorId != "" {
		result = result.Where("major_id = ?", MajorId)
	}

	if DepartmentId != "" {
		result = result.Where("department_id = ?", DepartmentId)
	}

	if TeamId != "" {
		result = result.Where("team_id = ?", TeamId)
	}

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&applications)

	data = GetListWithPagination(&applications, c, total)

	result.Find(&applications)
	return
}

func GetApplicationById(id interface{}) (n Application, err error) {
	err = db.First(&n, id).Error
	return
}

func (m *Application) Create() error {
	return db.Create(m).Error
}

func (m *Application) Delete() error {
	return db.Delete(m).Error
}

func (m *Application) Updates(n *Application) {
	db.Model(&Application{}).Where("id = ?", n.ID).Updates(n)
	db.Preload("Department").Preload("College").Preload("Major").Preload("Team").First(m, m.ID)
}
