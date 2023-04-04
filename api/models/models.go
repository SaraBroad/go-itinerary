package models

import (
	"time"

	"gorm.io/gorm"
)

type Itinerary struct {
	gorm.Model

	ID        string `gorm:"primary_key"`
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Locations []*Location      //1:M `gorm:"foreignKey:LocationID"`
	Items     []*ItineraryItem //1:M
}

type ItineraryItem struct {
	gorm.Model
	ID           string `gorm:"primary_key"`
	Name         string `json:"name"`
	Location     Location
	TimeAllotted string
	// Price        float64 //TODO:remove
	Price       *Price
	Category    *Category
	DayNumber   *DayNumber
	ItineraryID string
}

type Location struct {
	gorm.Model
	ID              string `gorm:"primary_key"`
	City            string
	Country         string
	ItineraryID     string
	ItineraryItemID string
}

type Price struct {
	gorm.Model
	ID              string `gorm:"primary_key"`
	IsFree          bool   //free or paid
	Amount          float64
	Currency        string
	ItineraryItemID string
}

type Category struct {
	gorm.Model
	ID              string `gorm:"primary_key"`
	Name            string
	ItineraryItemID string
}

type DayNumber struct {
	gorm.Model
	ID              string `gorm:"primary_key"`
	Num             int64
	ItineraryItemID string
}
