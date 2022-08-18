package orm

import (
	"go-todo/db"
	"go-todo/model"
)

type task struct{}

var Task task

const (
	DoingStatusTask  = "doing"
	DoneStatusTask   = "done"
	RejectStatusTask = "reject"
)

func (task) List() (task []model.Task, err error) {
	err = db.GetDB().Model(&task).Find(&task).Error
	return
}

func (task) Create(task *model.Task) (err error) {
	err = db.GetDB().Model(&task).Create(&task).Error

	return
}

func (task) GetById(id int) (task *model.Task, err error) {

	err = db.GetDB().Model(&task).Where("id = ?", id).First(&task).Error

	return
}

func (task) DeleteById(id int) (err error) {
	err = db.GetDB().Model(&model.Task{}).Where("id = ?", id).Delete(&model.Task{}).Error
	return
}

func (task) DoneTaskById(id int) (err error) {
	err = db.GetDB().Model(&model.Task{}).Where("id = ?", id).Update("status", DoneStatusTask).Error
	return
}

func (task) RejectTaskById(id int) (err error) {
	err = db.GetDB().Model(&model.Task{}).Where("id = ?", id).Update("status", RejectStatusTask).Error
	return
}

func (task) GetTaskByStatusOrTitle() (task []model.Task, err error) {
	title := "create CRUD 1"
	status := "doing"
	err = db.GetDB().Model(&model.Task{}).Where("title = ?", title).Or("status = ?", status).Find(&task).Error
	return
}
