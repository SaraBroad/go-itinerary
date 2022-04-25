package repository

import "github.com/SaraBroad/go-itinerary/models"

// use gorm
// has one day
// has one city
// has one category

type Config struct{}

func NewItinerary() {}

func CreateNewItem(item *models.Item) (*models.Item, error) {
	return &models.Item{}, nil
}
