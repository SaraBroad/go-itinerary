package service

import (
	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/SaraBroad/go-itinerary/api/repository"
)

// service layer - business logic

type database interface {
	CreateNewItem(item *models.Item) (*models.Item, error)
}

func calculateCategoryCost() (float64, error) {
	return 0, nil
}

func AddItem() {
	var db *repository.Database
	_, _ = db.CreateNewItem(&models.Item{})
}
