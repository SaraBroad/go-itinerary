package repository

import (
	"errors"
	"fmt"

	"github.com/SaraBroad/go-itinerary/api/models"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

// thinking about splitting this out from here - clean architecture
type ItineraryStore interface {
	CreateNewItinerary(itinerary *models.Itinerary) (*models.Itinerary, error)
	// add other methods
}

type DB struct {
	// define DB connection - DB *gorm.DB
}

// to here-ish

type Database struct {
	DB *gorm.DB
}

func NewItinerary(db *gorm.DB) Database {
	return Database{db}
}

func (db *Database) CreateNewItinerary(itinerary *models.Itinerary) (*models.Itinerary, error) {
	var newItinerary *models.Itinerary

	err := db.DB.Create(&newItinerary)
	if err != nil {
		return &models.Itinerary{}, errors.New("create new itinerary error")
	}

	return newItinerary, nil
}

func (db *Database) CreateNewItineraryItem(item *models.ItineraryItem) (*models.ItineraryItem, error) {
	fmt.Println("item", item)

	err := db.DB.Create(&item)
	if err != nil {
		fmt.Println("CreateNewItem", err)
		return &models.ItineraryItem{}, errors.New("create new items error")
	}
	return item, nil
}

func (db *Database) GetItinerary(id string) (*models.Itinerary, error) {
	return &models.Itinerary{}, nil
}

func (db *Database) GetItineraryItemById(id string) (*models.ItineraryItem, error) {
	if id == "" {
		return nil, errors.New("ID can't be nil")
	}
	var item *models.ItineraryItem
	if err := db.DB.Where("id = ?", id).First(&item); err != nil {
		return nil, errors.New("Can't find this item")
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
