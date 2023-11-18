package auth

import (
	"github.com/SaraBroad/go-itinerary/api/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Username    string `gorm:"size:255;not null;unique" json:"username"`
	Email       string `gorm:"size:255;not null;unique" json:"email"`
	Password    string `json:"-" binding:"required"`
	Itineraries []entity.Itinerary
}

func (user *User) CreateUserRecord() error {
	return nil
}
