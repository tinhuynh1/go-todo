package model

import (
	"time"
)

type Task struct {
	Id        int
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
