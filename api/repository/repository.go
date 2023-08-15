package repository

import (
	"errors"
	"fmt"

	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// thinking about splitting this out from here - clean architecture
type ItineraryStore interface {
	// CreateNewItinerary(itinerary *models.Itinerary) (*models.Itinerary, error)
	// add other methods
}

// to here-ish

type Database struct {
	DB *gorm.DB
}

func NewItinerary(db *gorm.DB) Database {
	return Database{db}
}

func (db *Database) CreateNewItinerary(itinerary *models.Itinerary) (*models.Itinerary, error) {
	itinerary.ID = uuid.New()
	fmt.Println("item.ID", itinerary.ID)
	err := db.DB.Create(&itinerary)
	if err != nil {
		return &models.Itinerary{}, errors.New("create itinerary error")
	}
	fmt.Println(itinerary.ID)
	return itinerary, nil
}

func (db *Database) FetchItinerary(id string) (*models.Itinerary, error) {
	fmt.Println("repo id", id)
	var itinerary *models.Itinerary
	// if err := db.DB.First(&itinerary, "id = ?", id); err != nil {

	// 	// if err := db.DB.First(&itinerary, "id = ?", id); err != nil {
	// 	fmt.Println("err", err.Error)
	// 	return &models.Itinerary{}, errors.New("fetch itinerary by id error")
	// }
	err := db.DB.First(&itinerary, "id = ?", id)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return nil, errors.New("error")
	}
	fmt.Println("here")
	fmt.Println("itinerary repo", itinerary)
	return itinerary, nil
}

func (db *Database) CreateNewItineraryItem(itineraryId string, item *models.ItineraryItem) (*models.ItineraryItem, error) {
	// get itinerary by id
	// create item within itinerary

	err := db.DB.Create(&item)
	if err != nil {
		fmt.Println("CreateNewItem", err)
		return &models.ItineraryItem{}, errors.New("create new items error")
	}
	return item, nil
}

func (db *Database) GetItineraryItemById(id string) (*models.ItineraryItem, error) {
	if id == "" {
		return nil, errors.New("id can't be nil")
	}
	var item *models.ItineraryItem
	if err := db.DB.Where("id = ?", id).First(&item); err != nil {
		return nil, errors.New("can't find this item")
	}

	return item, nil
}

func (db *Database) GetAllItinraryItems() ([]*models.ItineraryItem, error) {
	var items []*models.ItineraryItem
	if err := db.DB.Find(&items); err != nil {
		fmt.Println("GetAllItems err", err)
		return nil, errors.New("get all items error")
	}

	return items, nil
}

func (db *Database) GetItineraryItemByID(itemID string) (*models.ItineraryItem, error) {
	return &models.ItineraryItem{}, nil
}

func (db *Database) UpdateItinerary(id string, item models.ItineraryItem) (*models.ItineraryItem, error) {
	if id == "" {
		return nil, errors.New("ID can't be nil")
	}

	// for k := 0; k < len(items); k++ {
	// 	if item.itinerary[k].ID == id {
	// 		item.itinerary[k].Field = item.Field
	// 		return nil
	// 	}
	// }
	db.DB.Model(&item).Where("id = ?", id).Update("name", item.Name)

	return nil, nil
}

func (db *Database) UpdateItineraryItem(id string, item models.ItineraryItem) (*models.ItineraryItem, error) {
	if id == "" {
		return nil, errors.New("ID can't be nil")
	}
	db.DB.Model(&item).Where("id = ?", id).Update("name", item.Name)

	return nil, nil
}

func (db *Database) DeleteItinerary(id string) error {
	return nil
}

func (db *Database) DeleteItineraryItem(id string) error {
	fmt.Println("id repo", id)
	if id == "" {
		return errors.New("ID can't be nil")
	}
	var items *models.ItineraryItem
	if err := db.DB.Where("id = ?", id).Delete(&items); err != nil {
		return errors.New("Cannot delete")
	}
	return nil
}

func (db *Database) CreateNewLocation(loc *models.Location) (*models.Location, error) {
	i := &models.Location{
		// ID: uuid.NewString(),
	}

	return i, nil
}
