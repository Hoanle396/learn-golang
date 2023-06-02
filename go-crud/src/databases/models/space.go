package models

import "gorm.io/gorm"

type Space struct {
	gorm.Model

	Name      string  `json:"name"`
	ProjectID uint    `json:"project_id"`
	Project   Project `json:"project"`
	Tasks     []Task  `json:"tasks"`
}
