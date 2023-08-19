package seed

import "fmt"

// import (
// 	"fmt"
// 	"time"

// 	"github.com/SaraBroad/go-itinerary/api/models"
// 	"github.com/google/uuid"
// 	"github.com/pkg/errors"
// 	"gorm.io/gorm"
// )

// type Database struct {
// 	DB *gorm.DB
// }

func hello() {
	fmt.Println("hi")
}

// func (db *Database) Seed() (*models.Itinerary, error) {
// func (db *Database) Seed(itinerary *models.Itinerary) (*models.Itinerary, error) {
// 	i := &models.Itinerary{
// 		ID:        uuid.NewString(),
// 		StartDate: time.Now(),
// 		EndDate:   time.Now(),
// 		Locations: []*models.Location{{
// 			City:    "London",
// 			Country: "England",
// 		}},
// 		Items: []*models.ItineraryItem{{
// 			Name: "British Museum",
// 			Location: models.Location{
// 				City: "London",
// 			},
// 			Price: &models.Price{
// 				IsFree:   true,
// 				Amount:   0.00,
// 				Currency: "",
// 				// ItineraryItemID: "",
// 			},
// 			Category: &models.Category{
// 				Name: "activity",
// 			},
// 			DayNumber: &models.DayNumber{
// 				Num: 3,
// 			},
// 		}},
// 	}
// 	fmt.Println(i)
// 	err := db.DB.Save(&i)
// 	if err != nil {
// 		return nil, errors.Wrap(err.Error, "problem saving seed")
// 	}
// 	return i, nil
// // }
