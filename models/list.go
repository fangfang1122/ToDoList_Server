package models

type List struct {
	Model
	Name   string `json:"name" gorm:"not null"`
	UserId uint   `json:"user_id" gorm:"not null,index"`
	User   *User  `json:"user,omitempty" `
}

func GetAllList(UserID uint) (list []List) {

	result := db.Model(&List{}).Where("user_id = ?", UserID).Omit("User")

	result.Find(&list)
	return
}

func GetListById(id interface{}) (n List, err error) {
	err = db.First(&n, id).Error
	return
}

func (f *List) Create() error {
	return db.Create(f).Error
}

func (f *List) Update(n *List) {
	db.Model(&List{}).Where("id = ?", f.ID).Updates(n)
	db.Model(&List{}).Omit("User").First(f, f.ID)
}

func (f *List) Delete() error {
	return db.Delete(f).Error
}
