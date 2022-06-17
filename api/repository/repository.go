package repository

import (
	"errors"
	"fmt"

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

func (db *Database) GetItem(id string) (*models.Item, error) {
	if id == "" {
		return nil, errors.New("ID can't be nil")
	}

	return nil, nil
}

func (db *Database) GetAllItems() ([]*models.Item, error) {
	var items []*models.Item
	if err := db.DB.Find(&items); err != nil {
		fmt.Println("GetAllItems err", err)
		return nil, errors.New("get all items error")
	}

	return items, nil
}

func (db *Database) CreateNewItem(item *models.Item) (*models.Item, error) {
	fmt.Println("item", item)

	err := db.DB.Create(&item)
	if err != nil {
		fmt.Println("CreateNewItem", err)
		return &models.Item{}, errors.New("create new items error")
	}
	return item, nil
}

func (db *Database) UpdateItem(id string, item models.Item) (*models.Item, error) {
	if id == "" {
		return nil, errors.New("ID can't be nil")
	}
	db.DB.Model(&item).Where("id = ?", id).Update("name", item.Name)

	return nil, nil
}

func (db *Database) DeleteItem(id string) error {
	if id == "" {
		return errors.New("ID can't be nil")
	}
	if err := db.DB.Where("id = ?", id).Delete(id); err != nil {
		return errors.New("Cannot delete")
	}
	return nil
}
