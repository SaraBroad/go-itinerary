package service

import (
	"fmt"

	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/SaraBroad/go-itinerary/api/repository"
)

// service layer - business logic

type ItemService struct {
	itemRepo *repository.Database
}

func NewItemService(iRepo repository.Database) *ItemService {
	return &ItemService{
		itemRepo: &iRepo,
	}
}

type database interface {
	CreateNewItem(item *models.Item) (*models.Item, error)
}

func calculateCategoryCost() (float64, error) {
	return 0, nil
}

func (i *ItemService) AddItem(item *models.Item) error {
	_, err := i.itemRepo.CreateNewItem(item)
	if err != nil {
		fmt.Println("service error")
	}
	return nil
}
