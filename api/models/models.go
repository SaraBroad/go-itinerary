package models

import (
	"time"

	"gorm.io/gorm"
)

type Itinerary struct {
	gorm.Model
	ItineraryID  string `json:"itinerary_id" gorm:"primary_key"`
	StartDate    time.Time
	EndDate      time.Time
	Destinations []*Location      //1:M `gorm:"foreignKey:LocationID"`
	Items        []*ItineraryItem //1:M
}

type ItineraryItem struct {
	gorm.Model

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

	IsFree          bool //free or paid
	Amount          float64
	Currency        string
	ItineraryItemID string
}

type Category struct {
	gorm.Model

	Name            string
	ItineraryItemID string
}

type DayNumber struct {
	gorm.Model

	Num             int64
	ItineraryItemID string
}
