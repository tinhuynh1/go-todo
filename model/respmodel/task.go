package respmodel

import (
	"go-todo/model"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func NewTask(task *model.Task) Task {
	resp := Task{
		Id:        task.Id,
		Title:     task.Title,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: task.DeletedAt,
	}
	return resp
}

func NewListTask(tasks []model.Task) []Task {
	resp := make([]Task, len(tasks))
	for i, task := range tasks {
		resp[i] = NewTask(&task)
	}
	return resp
}
