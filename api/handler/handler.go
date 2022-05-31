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
	service service.ItemService
}

func NewItemHandler(is service.ItemService) *ItemHandler {
	fmt.Println("is", is)
	return &ItemHandler{service: is}
}

func (h *ItemHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	fmt.Print("HELLO")
	fmt.Println("w", w)
	fmt.Println("h", h.service)
	// vars := mux.Vars(r)
	fmt.Println("r", r)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("body", body)
	var item models.Item
	json.Unmarshal(body, &item)
	// z, _ := h.DB.CreateNewItem(&item)
	// i, err := repo.CreateNewItem(&models.Item{})
	err = h.service.AddItem(&item)
	fmt.Println(item)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Successfully Updated"))
}
