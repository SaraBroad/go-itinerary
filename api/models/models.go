package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	// ItemID    string
	Name      string
	Price     float64
	Category  Category
	DayNumber DayNumber
}

type Category struct {
	gorm.Model

	CategoryID string
	Name       string
	ItemID     string
}

type DayNumber struct {
	gorm.Model

	DayID  string
	DayNum int64
	ItemID string
}
