package models

type Department struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
}

func GetDepartmentList() (list []Department) {

	result := db.Model(&College{})

	result.Find(&list)
	return
}

func GetDepartmentById(id interface{}) (n Department, err error) {
	err = db.First(&n, id).Error
	return
}

func (m *Department) Create() error {
	return db.Create(m).Error
}

func (m *Department) Updates(n *Department) {
	db.Model(&Department{}).Where("id = ?", m.ID).Updates(n)
	db.Model(&Department{}).First(m, m.ID)
}

func (m *Department) Delete() error {
	return db.Delete(m).Error
}
