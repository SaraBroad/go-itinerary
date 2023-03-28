package repository

import (
	"fmt"
	"os"

	"github.com/SaraBroad/go-itinerary/api/models"
	// "github.com/jinzhu/gorm"
	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dburl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("PORT"),
	)

	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
	// db, err := gorm.Open("postgres", dburl)
	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()
	db.AutoMigrate(&models.Itinerary{}, &models.ItineraryItem{}, &models.Location{}, &models.Category{}, &models.Price{}, &models.DayNumber{})

	return db
}
