package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"passowrd" binding:"required"`
}

func (user *User) CreateUserRecord() error {
	return nil
}
