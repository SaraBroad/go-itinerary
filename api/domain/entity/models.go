package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Itinerary struct {
	gorm.Model

	ID        uuid.UUID `gorm:"primaryKey;autoIncrement" json:"itinerary_id"`
	Name      string    `json:"itinerary_name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	//1:M `gorm:"foreignKey:LocationID"
	Locations []*Location `json:"locations"`
	//1:M
	Items []*ItineraryItem `json:"items"`
}

type ItineraryItem struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primary_key;autoIncrement" json:"itinerary_item_id"`
	Name         string    `json:"item_name"`
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
	ID              uuid.UUID `gorm:"primary_key;autoIncrement" json:"location_id"`
	City            string    `json:"city"`
	Country         string
	ItineraryID     uuid.UUID
	ItineraryItemID uuid.UUID
}

type Price struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primary_key;autoIncrement" json:"price_id"`
	IsFree          bool      //free or paid
	Amount          float64
	Currency        string
	ItineraryItemID uuid.UUID
}

type Category struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primary_key" json:"category_id"`
	Name            string
	ItineraryItemID uuid.UUID
}

type DayNumber struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primary_key" json:"day_num_id"`
	Num             int64
	ItineraryItemID uuid.UUID
}
