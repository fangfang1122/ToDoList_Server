package models

type Major struct {
	ID        uint     `json:"id" gorm:"primary_key"`
	Name      string   `json:"name" gorm:"not null"`
	CollegeId uint     `json:"college_id" gorm:"not null"`
	College   *College `json:"college,omitempty" `
}

func GetMajorList(collegeId interface{}) (majors []Major) {

	result := db.Model(&Major{}).Preload("College")

	if collegeId != "" {
		result = result.Where("college_id = ?", collegeId)
	}

	result.Find(&majors)
	return
}

func GetMajorById(id interface{}) (n Major, err error) {
	err = db.First(&n, id).Error
	return
}

func (m *Major) Create() error {
	return db.Create(m).Error
}

func (m *Major) Updates(n *Major) {
	db.Model(&Major{}).Where("id = ?", m.ID).Updates(n)
	db.Model(&Major{}).First(m, m.ID)
}

func (m *Major) Delete() error {
	return db.Delete(m).Error
}
