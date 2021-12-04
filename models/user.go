package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username"`
	Openid string `json:"-" gorm:"uniqueIndex"`
	Avatar string `json:"avatar"`
}

func GetUserById(id interface{}) (n User, err error) {
	err = db.First(&n, id).Error
	return
}

func GetUserByOpenId(openId string) (n User, err error) {
	result :=db.Where("openid=?",openId).First(&n)
	if result.RowsAffected ==0{
		n = User{
			Username: "",
			Avatar: "",
			Openid: openId,
		}
		err = n.Create()
	}
	return
}

func (f *User) Create() error {
	return db.Create(f).Error
}

func (f *User) Update(n *User) {
	db.Model(&User{}).Where("id = ?", f.ID).Updates(n)
	db.Model(&User{}).First(f, f.ID)
}

func (f *User) Delete() error {
	return db.Delete(f).Error
}
