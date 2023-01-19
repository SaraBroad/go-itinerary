package repository

import (
	"fmt"

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
	Timezone string
}

func InitDatabase(auth *Auth) *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", auth.Host, auth.Username, auth.Password, auth.DBName, auth.Port, auth.SSL, auth.Timezone)
	//db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	db, err := gorm.Open(dbUrl)
	if err != nil {
		// log.Fatalln(err)
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&models.Itinerary{}, &models.ItineraryItem{}, &models.DayNumber{}, &models.Category{}, &models.Price{}, &models.Location{})
	return db
}
