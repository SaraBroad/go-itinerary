package repository

import (
	"fmt"
	"log"

	"github.com/SaraBroad/go-itinerary/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Item{}, &models.DayNumber{}, &models.Category{})

	i := &models.Item{
		Name:  "Hyatt",
		Price: 199.99,
		Category: models.Category{
			Name: "Accomodations",
		},
		DayNumber: models.DayNumber{
			DayNum: 1,
		},
	}
	db.Create(i)
	return db
}
