package repository

import (
	"github.com/SaraBroad/go-itinerary/api/models"
	"gorm.io/gorm"
)

// use gorm
// has one day
// has one city
// has one category

type Database struct {
	DB gorm.DB
}

func NewItinerary() Database {
	return Database{}
}

func (db Database) CreateNewItem(item *models.Item) (*models.Item, error) {
	return &models.Item{}, nil
}
