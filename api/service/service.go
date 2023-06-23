package service

import (
	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/SaraBroad/go-itinerary/api/repository"
	"github.com/pkg/errors"
)

// service layer - business logic

type ItineraryService struct {
	itineraryRepo *repository.Database
}

func NewItineraryService(iRepo repository.Database) *ItineraryService {
	return &ItineraryService{
		itineraryRepo: &iRepo,
	}
}

type database interface {
	CreateNewItineraryItem(item *models.ItineraryItem) (*models.ItineraryItem, error)
	GetItineraryItemById(id string) (*models.ItineraryItem, error)
	UpdateItineraryItem(id string, item models.ItineraryItem) (*models.ItineraryItem, error)
	DeleteItineraryItem(id string) error
}

// func calculateCategoryCost(catCost *models.ItineraryItem) (float64, error) {
// 	price := catCost.Price
// 	var cost float64
// 	switch catCost.Category.Name {
// 	case "accomodations":
// 		cost += price
// 	case "activity":
// 		cost += price
// 	case "transportation":
// 		cost += price
// 	case "shopping":
// 		cost += price
// 	case "food":
// 		cost += price
// 	default:
// 		fmt.Println("unknown category")
// 	}
// 	return cost, nil
// }

func calculateTotalCost() (float64, error) {
	return 0, nil
}

func (i *ItineraryService) AddItinerary(itinerary *models.Itinerary) (*models.Itinerary, error) {
	newItinerary, err := i.itineraryRepo.CreateNewItinerary(itinerary)
	if err != nil {
		return nil, errors.Wrap(err, "add itinerary error")
	}

	return newItinerary, nil
}

func (i *ItineraryService) GetItineraryById(id string) (*models.Itinerary, error) {
	itinerary, err := i.itineraryRepo.FetchItinerary(id)
	if err != nil {
		return nil, errors.Wrap(err, "get itinerary by id error")
	}

	return itinerary, nil
}

// func (i *ItineraryService) AddItem(item *models.ItineraryItem) (*models.ItineraryItem, error) {

// 	newItem, err := i.itineraryRepo.CreateNewItineraryItem(item)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "add item service error")
// 	}
// 	return newItem, nil
// }

// func (i *ItineraryService) FindItemById(id string) (*models.ItineraryItem, error) {
// 	item, err := i.itineraryRepo.GetItineraryItemById(id)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "find item by id service error")
// 	}
// 	return item, nil
// }

// func (i *ItineraryService) Update(id string, item *models.ItineraryItem) (*models.ItineraryItem, error) {
// 	item, err := i.itineraryRepo.UpdateItinerary(id, *item)
// 	if err != nil {
// 		fmt.Println("Update item service error")
// 	}
// 	return item, nil
// }

// func (i *ItineraryService) RemoveItem(id string) error {
// 	fmt.Println("id service", id)
// 	if err := i.itineraryRepo.DeleteItineraryItem(id); err != nil {
// 		fmt.Println("error deleting item in service", err)
// 	}
// 	return nil
// }
