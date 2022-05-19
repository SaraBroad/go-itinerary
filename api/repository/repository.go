package repository

import (
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

func (db *Database) CreateNewItem(item *models.Item) (*models.Item, error) {
	fmt.Println("item", item)

	item = &models.Item{
		Name:  item.Name,
		Price: item.Price,
		Category: models.Category{
			Name: item.Category.Name,
		},
		DayNumber: models.DayNumber{
			DayNum: item.DayNumber.DayNum,
		},
	}

	_ = db.DB.Create(&item)
	return item, nil
}
