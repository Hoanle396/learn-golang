package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Name     string  `json:"name"`
	OwnerID  uint    `json:"owner_id"`
	Owner    User    `json:"owner"`
	LeaderID uint    `json:"leader_id"`
	Leader   User    `json:"leader"`
	Members  []User  `json:"members" gorm:"many2many:user_projects;"`
	Spaces   []Space `json:"spaces"`
}
