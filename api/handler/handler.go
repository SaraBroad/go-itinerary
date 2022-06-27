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

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("body", body)
	var item models.Item
	json.Unmarshal(body, &item)

	_, err = h.service.AddItem(&item)
	fmt.Println(item)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Successfully Added"))
}

func (h *ItemHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		fmt.Println("delete item handler error")
	}

	err = h.service.RemoveItem(item.ID)
}

func (h *ItemHandler) GetItemByItemId(w http.ResponseWriter, r *http.Request) {}

func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {}
