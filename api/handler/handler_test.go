package handler

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/SaraBroad/go-itinerary/api/models"
)

func Test_AddItinerary(t *testing.T) {
	input := `{
		"id": "123",
		"itinerary_name": "Summer Fun",
		"start_date": "2023-07-01T12:00:00.000Z",
		"end_date": "2023-07-10T14:00:20.000Z"
	}`
	itinerary := &models.Itinerary{}
	err := json.Unmarshal([]byte(input), itinerary)
	if err != nil {
		return
	}
	fmt.Println("itinerary", itinerary)
}
