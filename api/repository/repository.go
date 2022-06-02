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

// func NewItinerary() Database {
// 	return Database{}
// }

func NewItinerary(db *gorm.DB) Database {
	return Database{db}
}

func (db *Database) GetItem(id string) (*models.Item, error) {
	return nil, nil
}

func (db *Database) GetAllItems() ([]*models.Item, error) {
	return nil, nil
}

func (db *Database) CreateNewItem(item *models.Item) (*models.Item, error) {
	fmt.Println("item", item)

	err := db.DB.Create(&item)
	if err != nil {
		fmt.Println("err", err)
		fmt.Println("Create New Item Error")
	}
	return item, nil
}

func (db *Database) UpdateItem(id string, item models.Item) (*models.Item, error) {
	db.DB.Model(&item).Where("id = ?", id).Update("name", item.Name)

	return nil, nil
}

func (db *Database) DeleteItem(id string) error {
	if id == "" {
		return errors.New("ID can't be nil")
	}
	return nil
}
