package service

import (
	"fmt"

	"github.com/SaraBroad/go-itinerary/api/domain/entity"
	"github.com/SaraBroad/go-itinerary/api/domain/repository"
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
	CreateNewItineraryItem(item *entity.ItineraryItem) (*entity.ItineraryItem, error)
	GetItineraryItemById(id string) (*entity.ItineraryItem, error)
	UpdateItinerary(id string, item entity.Itinerary) (*entity.ItineraryItem, error)
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

func (i *ItineraryService) AddItinerary(itinerary *entity.Itinerary) (*entity.Itinerary, error) {
	newItinerary, err := i.itineraryRepo.CreateNewItinerary(itinerary)
	if err != nil {
		return nil, errors.Wrap(err, "add itinerary error")
	}

	return newItinerary, nil
}

func (i *ItineraryService) GetItineraryById(id string) (*entity.Itinerary, error) {
	fmt.Println(id)
	itinerary, err := i.itineraryRepo.FetchItinerary(id)
	if err != nil {
		return nil, errors.Wrap(err, "get itinerary by error")
	}
	fmt.Println("itinerary", itinerary)
	return itinerary, nil
}

func (i *ItineraryService) GetAllItineraries() ([]*entity.Itinerary, error) {
	fmt.Println("service")
	itineraries, err := i.itineraryRepo.FetchAllItineraries()
	if err != nil {
		return nil, errors.Wrap(err, "get itineraries by error")
	}
	fmt.Println("itinerary", itineraries)
	return itineraries, nil
}

func (i *ItineraryService) AddItem(itineraryId string, item *entity.ItineraryItem) (*entity.ItineraryItem, error) {
	fmt.Println("itinerary id service", itineraryId)
	fmt.Println("service item", item)
	if item.Name == "" {
		return nil, errors.New("your item requires a name")
	}

	newItem, err := i.itineraryRepo.CreateNewItineraryItem(itineraryId, item)
	if err != nil {
		return nil, errors.Wrap(err, "add item service error")
	}
	return newItem, nil
}

func (i *ItineraryService) FindItemById(id string) (*entity.ItineraryItem, error) {
	item, err := i.itineraryRepo.GetItineraryItemById(id)
	if err != nil {
		return nil, errors.Wrap(err, "find item by id service error")
	}
	return item, nil
}

func (i *ItineraryService) UpdateItineraryById(id string, item *entity.Itinerary) (*entity.Itinerary, error) {
	itinerary, err := i.itineraryRepo.UpdateItinerary(id, *item)
	if err != nil {
		fmt.Println("Update item service error")
	}
	return itinerary, nil
}

// func (i *ItineraryService) UpdateItineraryItemById(id string, item *entity.ItineraryItem) (*entity.ItineraryItem, error) {
// 	item, err := i.itineraryRepo.UpdateItinerary(id, *item)
// 	if err != nil {
// 		fmt.Println("Update item service error")
// 	}
// 	return item, nil
// }

func (i *ItineraryService) RemoveItem(id string) error {
	fmt.Println("id service", id)
	if err := i.itineraryRepo.DeleteItineraryItem(id); err != nil {
		fmt.Println("error deleting item in service", err)
	}
	return nil
}
