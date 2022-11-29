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
	CreateNewItineraryItem(item *models.ItineraryItem) (*models.ItineraryItem, error)
	GetItineraryItemById(id string) (*models.ItineraryItem, error)
	UpdateItineraryItem(id string, item models.ItineraryItem) (*models.ItineraryItem, error)
	DeleteItineraryItem(id string) error
}

func calculateCategoryCost(catCost *models.ItineraryItem) (float64, error) {
	price := catCost.Price
	var cost float64
	switch catCost.Category.Name {
	case "accomodations":
		cost += price
	case "activity":
		cost += price
	case "transportation":
		cost += price
	case "shopping":
		cost += price
	case "food":
		cost += price
	default:
		fmt.Println("unknown category")
	}
	return cost, nil
}

func calculateTotalCost() (float64, error) {
	return 0, nil
}

func (i *ItemService) AddItem(item *models.ItineraryItem) (*models.ItineraryItem, error) {

	newItem, err := i.itemRepo.CreateNewItineraryItem(item)
	if err != nil {
		fmt.Println("add item service error")
	}
	return newItem, nil
}

func (i *ItemService) FindItemById(id string) (*models.ItineraryItem, error) {
	item, err := i.itemRepo.GetItineraryItemById(id)
	if err != nil {
		fmt.Println("find item by id service error")
	}
	return item, nil
}

func (i *ItemService) Update(id string, item *models.ItineraryItem) (*models.ItineraryItem, error) {
	item, err := i.itemRepo.UpdateItem(id, *item)
	if err != nil {
		fmt.Println("Update item service error")
	}
	return item, nil
}

func (i *ItemService) RemoveItem(id string) error {
	fmt.Println("id service", id)
	if err := i.itemRepo.DeleteItineraryItem(id); err != nil {
		fmt.Println("error deleting item in service", err)
	}
	return nil
}
