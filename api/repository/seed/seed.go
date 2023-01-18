package seed

import (
	"fmt"
	"time"

	"github.com/SaraBroad/go-itinerary/api/models"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) Seed(itinerary *models.Itinerary) (*models.Itinerary, error) {
	i := &models.Itinerary{
		StartDate: time.Now(),
		EndDate:   time.Now(),
		Destinations: []*models.Location{{
			ID:      "",
			City:    "",
			Country: "",
		}},
		Items: []*models.ItineraryItem{{
			Name: "",
			Location: models.Location{
				City: "",
			},
			Pce: &models.Price{
				ID:              "",
				IsPaid:          true,
				Amount:          0.00,
				Currency:        "",
				ItineraryItemID: "",
			},
			Category:  &models.Category{},
			DayNumber: &models.DayNumber{},
		}},
	}
	fmt.Println(i)

	return i, nil
}
