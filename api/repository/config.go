package repository

import (
	"log"

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
}

func InitDatabase(auth Auth) *gorm.DB {
	dbUrl := ""
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
