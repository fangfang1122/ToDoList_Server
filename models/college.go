package models

type College struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"unique"`
}

func GetCollegeList() (colleges []College) {

	result := db.Model(&College{})

	result.Find(&colleges)
	return
}

func GetCollegeById(id interface{}) (n College, err error) {
	err = db.First(&n, id).Error
	return
}

func (m *College) Create() error {
	return db.Create(m).Error
}

func (college *College) Updates(n *College) {
	db.Model(&College{}).Where("id = ?", college.ID).Updates(n)
	db.Model(&College{}).First(college, college.ID)
}

func (m *College) Delete() error {
	return db.Delete(m).Error
}
