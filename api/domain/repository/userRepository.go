package repository

import (
	"github.com/SaraBroad/go-itinerary/api/domain/entity"
	"gorm.io/gorm"
)

// type Database struct {
// 	DB *gorm.DB
// }

func NewUser(db *gorm.DB) Database {
	return Database{db}
}

func (db *Database) CreateUserRecord(user *entity.User) (*entity.User, error) {
	err := db.DB.Create(&user).Error
	if err != nil {
		return &entity.User{}, err
	}
	return user, nil
}

func (db *Database) BeforeSave() error {
	// GenerateFromPassword
	return nil
}

func (db *Database) FindUserById(id string) (entity.User, error) {
	var user entity.User
	// don't think i need to find itineraries
	err := db.DB.Preload("Itineraries").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return entity.User{}, nil
	}
	return user, nil
}
