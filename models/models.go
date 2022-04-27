package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	ID    string  `json:"item_id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Category
	DayNumber
}

type Category struct {
	gorm.Model

	ID   string `json:"category_id"`
	Name string `json:"name"`
}

type DayNumber struct {
	gorm.Model

	ID     string `json:"day_id"`
	DayNum int64  `json:"day_num"`
}
