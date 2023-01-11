package models

import (
	"time"

	"gorm.io/gorm"
)

type Itinerary struct {
	gorm.Model

	StartDate    time.Time
	EndDate      time.Time
	Destinations []*Location      //1:M
	Items        []*ItineraryItem //1:M
}

type ItineraryItem struct {
	gorm.Model

	ID           string
	Name         string
	Location     Location
	TimeAllotted string
	Price        float64 //TODO:remove
	Pce          *Price
	Category     *Category
	DayNumber    *DayNumber
	ItineraryID  string
}

type Location struct {
	gorm.Model

	ID              string
	City            string
	Country         string
	ItineraryID     string
	ItineraryItemID string
}

type Price struct {
	gorm.Model

	ID              string
	IsPaid          bool //free or paid
	Amount          float64
	Currency        string
	ItineraryItemID string
}

type Category struct {
	gorm.Model

	ID              string
	Name            string
	ItineraryItemID string
}

type DayNumber struct {
	gorm.Model

	ID              string
	Num             int64
	ItineraryItemID string
}
