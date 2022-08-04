package entity

import (
	"go-todo/model"
	"go-todo/orm"
)

func GetAllTask() (tasks []model.Task, err error) {
	tasks, err = orm.Task.List()
	if err != nil {
		return
	}
	return

}
func GetById(id int) (task *model.Task, err error) {
	task, err = orm.Task.GetById(id)
	if err != nil {
		return
	}
	return

}
func CreateTask(task model.Task) (err error) {
	task.Status = orm.DoingStatusTask
	err = orm.Task.Create(&task)
	if err != nil {
		return
	}
	return
}

func DeleteById(id int) (err error) {
	err = orm.Task.DeleteById(id)
	if err != nil {
		return
	}
	return

}

func DoneTaskById(id int) (err error) {
	err = orm.Task.DoneTaskById(id)
	if err != nil {
		return
	}
	return

}

func RejectTaskById(id int) (err error) {
	err = orm.Task.RejectTaskById(id)
	if err != nil {
		return
	}

	return

}
