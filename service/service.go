package service

import "github.com/SaraBroad/go-itinerary/models"

// service layer - business logic

type database interface {
	CreateNewItem(item *models.Item) (*models.Item, error)
}

func calculateCategoryCost() (float64, error) {
	return 0, nil
}
