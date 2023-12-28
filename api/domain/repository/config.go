package repository

import (
	"fmt"
	"os"

	// "github.com/jinzhu/gorm"

	models "github.com/SaraBroad/go-itinerary/api/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// _ = godotenv.Load("")

	dburl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("PORT"),
	)
	fmt.Println("dburl", dburl)
	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
	// db, err := gorm.Open("postgres", dburl)
	if err != nil {
		panic(err.Error())
	}

	// defer db.Close()
	db.AutoMigrate(&models.Itinerary{}, &models.ItineraryItem{}, &models.Location{}, &models.Category{}, &models.Price{}, &models.DayNumber{})
	t, _ := db.Migrator().GetTables()
	fmt.Println("tables", t)
	return db
}

type Database struct {
	DB *gorm.DB
}
