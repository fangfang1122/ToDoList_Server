package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username         string `json:"username"`
	Openid           string `json:"-" gorm:"uniqueIndex"`
	Avatar           string `json:"avatar"`
	IsClickHeavy     bool   `json:"is_click_heavy" gorm:"default:'true'"`
	IsClickSound     bool   `json:"is_click_sound" gorm:"default:'true'"`
	CreateListAmount int    `json:"create_list_amount" gorm:"default:'0'"`
}

func GetUserById(id interface{}) (n User, err error) {
	err = db.First(&n, id).Error
	return
}

func GetUserByOpenId(openId string) (n User, err error) {
	result := db.Where("openid=?", openId).First(&n)
	if result.RowsAffected == 0 {
		n = User{
			Username:         "",
			Avatar:           "",
			Openid:           openId,
			IsClickHeavy:     true,
			IsClickSound:     true,
			CreateListAmount: 0,
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
