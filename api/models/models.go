package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	ItemID string  `gorm:"primaryKey"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Category
	DayNumber
}

type Category struct {
	gorm.Model

	CategoryID string `gorm:"primaryKey"`
	Name       string `json:"name"`
}

type DayNumber struct {
	gorm.Model

	ID     string `gorm:"primaryKey"`
	DayNum int64  `json:"day_num"`
}
