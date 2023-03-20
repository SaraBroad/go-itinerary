package repository

import (
	"fmt"
	"os"

	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	// "gorm.io/gorm"
)

type Auth struct {
	Username string
	Password string
	DBName   string
	Port     string
	Host     string
	URI      string
	SSL      string
	// Timezone string
}

// var DB *gorm.DB

// func InitDatabase(auth *Auth) *gorm.DB {
// 	fmt.Println("auth", auth)
// 	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", auth.Host, auth.Username, auth.Password, auth.DBName, auth.Port, auth.SSL, auth.Timezone)

// 	db, err := gorm.Open("postgres", dbUrl)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	db.AutoMigrate(&models.Itinerary{}, &models.ItineraryItem{}, &models.Location{}, &models.Category{}, &models.Price{}, &models.DayNumber{})
// 	return db
// }

func ConnectDatabase() *gorm.DB {
	dburl := fmt.Sprintf("host=%s username=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("HOST"),
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("PORT"),
		os.Getenv("SSL"),
	)

	db, err := gorm.Open("postgres", dburl)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	db.AutoMigrate(&models.Itinerary{}, &models.ItineraryItem{}, &models.Location{}, &models.Category{}, &models.Price{}, &models.DayNumber{})

	return db
}
