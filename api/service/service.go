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
	UpdateItem(id string, item models.Item) (*models.Item, error)
	DeleteItem(id string) error
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

func (i *ItemService) Update(id string, item *models.Item) (*models.Item, error) {
	item, err := i.itemRepo.UpdateItem(id, *item)
	if err != nil {
		fmt.Println("Update item service error")
	}
	return item, nil
}

func (i *ItemService) RemoveItem(id string) error {
	fmt.Println("id service", id)
	if err := i.itemRepo.DeleteItem(id); err != nil {
		fmt.Println("error deleting item in service", err)
	}
	return nil
}
