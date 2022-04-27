package repository

import (
	"github.com/SaraBroad/go-itinerary/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// use gorm
// has one day
// has one city
// has one category

type Config struct {
}

func Init() *gorm.DB {
	dbURL := ""
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	return db
}

func NewItinerary() {}

func CreateNewItem(item *models.Item) (*models.Item, error) {
	return &models.Item{}, nil
}
