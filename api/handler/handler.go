package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/SaraBroad/go-itinerary/api/service"
)

// http layer

type ItemHandler struct {
	// DB *repository.Database
	service *service.ItemService
}

func NewItemHandler(is service.ItemService) *ItemHandler {
	return &ItemHandler{service: &is}
}

func (h *ItemHandler) AddItem(w http.ResponseWriter, r *http.Request) {

	// func (h *HTTP) AddItem(w http.ResponseWriter, r *http.Request) {
	// db := repository.InitDatabase(&repository.Auth{
	// 	Username: os.Getenv("USERNAME"),
	// 	Host:     os.Getenv("HOST"),
	// 	Password: os.Getenv("PASSWORD"),
	// 	DBName:   os.Getenv("DBNAME"),
	// 	Port:     os.Getenv("PORT"),
	// 	SSL:      os.Getenv("SSL"),
	// 	Timezone: os.Getenv("TIMEZONE"),
	// })

	// repo := repository.NewItinerary(db)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	var item models.Item
	json.Unmarshal(body, &item)
	// z, _ := h.DB.CreateNewItem(&item)
	// i, err := repo.CreateNewItem(&models.Item{})
	err = h.service.AddItem(&item)
	fmt.Println(item)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
}
