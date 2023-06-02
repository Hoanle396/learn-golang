package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     time.Time `json:"due_date"`
	SpaceID     uint   `json:"space_id"`
	Space       Space  `json:"space"`
	UserID      uint   `json:"user_id"`
	User        User   `json:"user"`
}
