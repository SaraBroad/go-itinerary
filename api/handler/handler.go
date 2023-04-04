package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SaraBroad/go-itinerary/api/models"
	"github.com/SaraBroad/go-itinerary/api/service"
	"github.com/gorilla/mux"
)

// http layer

type ItineraryHandler struct {
	// DB *repository.Database
	service service.ItineraryService
}

func NewItemHandler(svc service.ItineraryService) *ItineraryHandler {
	return &ItineraryHandler{service: svc}
}

func (h *ItineraryHandler) AddItem(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("body", body)
	var item models.ItineraryItem
	json.Unmarshal(body, &item)

	_, err = h.service.AddItem(&item)
	fmt.Println(item)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Successfully Added"))
}

func (h *ItineraryHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("id", id)
	if err := h.service.RemoveItem(id); err != nil {
		fmt.Println("err", err)
	}
}

func (h *ItineraryHandler) GetItemByItemId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("id", id)
	var item *models.ItineraryItem
	if id == "" {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
	fmt.Println("item", item)
}

func (h *ItineraryHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {}

func indexByID(items []models.ItineraryItem, id string) int {
	for i := 0; i < len(items); i++ {
		if items[i].ItineraryID == id {
			return i
		}
	}
	return -1
}
