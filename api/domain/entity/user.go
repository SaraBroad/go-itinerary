package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Username    string `gorm:"size:255;not null;unique" json:"username"`
	Password    string `json:"-" binding:"required"`
	Itineraries []Itinerary
}

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
