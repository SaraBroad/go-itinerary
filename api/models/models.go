package models

import (
	"time"

	"gorm.io/gorm"
)

type Itinerary struct {
	gorm.Model

	startDate    time.Time
	endDate      time.Time
	destinations []Location      //1:M
	items        []ItineraryItem //1:M
}

type ItineraryItem struct {
	gorm.Model

	ID           string
	Name         string
	Location     Location //1:1
	timeAllotted string
	Price        float64   //TODO:remove
	P            Price     //1:1
	Category     Category  //M:1
	DayNumber    DayNumber //M:1
}

type Location struct {
	id      string
	city    string
	country string
}

type Price struct {
	id       string
	cost     bool //free or paid
	amount   float64
	currency string
}

type Category struct {
	gorm.Model

	CategoryID string
	Name       string
	ItemID     string
}

type DayNumber struct {
	gorm.Model

	DayID  string
	DayNum int64
	ItemID string
}
