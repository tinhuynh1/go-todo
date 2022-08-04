package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	Id        int
	Title     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
