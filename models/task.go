package models

import (
	"github.com/gin-gonic/gin"
)

type Task struct {
	Model
	Name         string     `json:"name" gorm:"not null"`
	UserId       uint       `json:"user_id" gorm:"not null,index"`
	ListId       uint       `json:"list_id" gorm:"not null"`
	List         *List      `json:"list,omitempty" `
	EndDate      TimeNormal `json:"end_date" gorm:"default:'null'"`
	Week         int        `json:"week"`
	IsFinished   bool       `json:"is_finished"`
	FinishedDate TimeNormal `json:"finished_date" gorm:"default:'null'""`
	Description  string     `json:"description"`
	Files        []File     `json:"files"`
	Photos       []Photo    `json:"photos"`
}

func GetTaskList(c *gin.Context, userId interface{}) (data *DataList) {
	var list []Task
	result := db.Model(&Task{})

	if c.Query("list_id") != "" {
		result = result.Where("list_id=?", c.Query("list_id"))
	}

	result = result.Where("user_id=?", userId).Preload("Files").Preload("Photos")

	var total int64

	result.Count(&total)

	result.Scopes(orderAndPaginate(c)).Find(&list)

	data = GetListWithPagination(&list, c, total)
	return
}

func GetAllTask(UserID uint, ListID interface{}) (list []Task) {

	result := db.Model(&Task{}).Where("user_id = ?", UserID).Where("list_id=?", ListID)

	result.Find(&list)
	return
}

func GetTaskById(id interface{}) (n Task, err error) {
	err = db.First(&n, id).Error
	return
}

func (f *Task) Create() error {
	return db.Create(f).Error
}

// 哭了啊，记一次踩坑记录，updates只会更新非零值的字段，所以bool值最好还是自己单独update
func (f *Task) CancelFinish() {
	db.Model(&Task{}).Where("id = ?", f.ID).Update("is_finished", false)
}

func (f *Task) Update(n *Task) {
	db.Model(&Task{}).Where("id = ?", f.ID).Updates(n)
	db.Model(&Task{}).Preload("Files").Preload("Photos").First(f, f.ID)
}

func (f *Task) Delete() error {
	return db.Delete(f).Error
}
