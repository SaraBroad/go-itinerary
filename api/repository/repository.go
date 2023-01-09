package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/SaraBroad/go-itinerary/api/models"
	"gorm.io/gorm"
)

// use gorm
// has one day
// has one city
// has one category

type Database struct {
	DB *gorm.DB
}

func NewItinerary(db *gorm.DB) Database {
	return Database{db}
}

func (db *Database) Seed(itinerary *models.Itinerary) (*models.Itinerary, error) {
	i := &models.Itinerary{
		StartDate: time.Now(),
		EndDate:   time.Now(),
		Destinations: []*models.Location{{
			ID:      "",
			City:    "",
			Country: "",
		}},
		Items: []*models.ItineraryItem{{
			Name: "",
			Location: models.Location{
				City: "",
			},
			Pce: &models.Price{
				ID:              "",
				IsCost:          true,
				Amount:          0.00,
				Currency:        "",
				ItineraryItemID: "",
			},
			Category:  &models.Category{},
			DayNumber: &models.DayNumber{},
		}},
	}
	fmt.Println(i)

	return i, nil
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

// List itinerary
func (db *Database) GetAllItinraryItems() ([]*models.ItineraryItem, error) {
	var items []*models.ItineraryItem
	if err := db.DB.Find(&items); err != nil {
		fmt.Println("GetAllItems err", err)
		return nil, errors.New("get all items error")
	}

	return items, nil
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

func (db *Database) UpdateItem(id string, item models.ItineraryItem) (*models.ItineraryItem, error) {
	if id == "" {
		return nil, errors.New("ID can't be nil")
	}
	db.DB.Model(&item).Where("id = ?", id).Update("name", item.Name)

	return nil, nil
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
