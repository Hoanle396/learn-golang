package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Verified bool   `json:"verified" gorm:"default:false"`
	Wallet   int64  `json:"wallet" gorm:"default:0"`
	ProjectsOwned    []Project `json:"project_owner" gorm:"foreignKey:OwnerID"`
    ProjectsLeaded   []Project `json:"project_leader" gorm:"foreignKey:LeaderID"`
    ProjectsInvolved []Project `json:"project_involved" gorm:"many2many:user_projects;"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}