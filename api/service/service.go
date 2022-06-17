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
	GetItemById(id string) (*models.Item, error)
}

func calculateCategoryCost() (float64, error) {
	return 0, nil
}

func (i *ItemService) AddItem(item *models.Item) (*models.Item, error) {
	newItem, err := i.itemRepo.CreateNewItem(item)
	if err != nil {
		fmt.Println("add item service error")
	}
	return newItem, nil
}

func (i *ItemService) FindItemById(id string) (*models.Item, error) {
	item, err := i.itemRepo.GetItemById(id)
	if err != nil {
		fmt.Println("find item by id service error")
	}
	return item, nil
}
