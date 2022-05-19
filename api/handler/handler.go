package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/SaraBroad/go-itinerary/api/repository"
)

// http layer

type HTTP struct{}

func AddItem(w http.ResponseWriter, r *http.Request) {
	db := repository.InitDatabase(&repository.Auth{
		Username: os.Getenv("USERNAME"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		Port:     os.Getenv("PORT"),
		SSL:      os.Getenv("SSL"),
		Timezone: os.Getenv("TIMEZONE"),
	})

	repo := repository.NewItinerary(db)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	var item models.Item
	json.Unmarshal(body, &item)
	i, err := repo.CreateNewItem(&models.Item{})
	fmt.Println(i)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
}
