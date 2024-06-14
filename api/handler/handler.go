package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SaraBroad/go-itinerary/api/domain/entity"
	"github.com/SaraBroad/go-itinerary/api/service"
	"github.com/gorilla/mux"
)

// http layer
type ItineraryHandler struct {
	service service.ItineraryService
}

func NewItemHandler(svc service.ItineraryService) *ItineraryHandler {
	return &ItineraryHandler{service: svc}
}

func (h *ItineraryHandler) AddItinerary(w http.ResponseWriter, r *http.Request) {
	var itinerary *entity.Itinerary
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
	fmt.Println("r.body", r.Body)
	err = json.Unmarshal(body, &itinerary)
	fmt.Println("handler", itinerary.Name)

	if err != nil {
		fmt.Println(errors.New("cannot unmarshal new itinerary"))
	}
	_, err = h.service.AddItinerary(itinerary)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Successfully added"))
}

func (h *ItineraryHandler) GetItinerary(w http.ResponseWriter, r *http.Request) {
	// json new encoder
	id := mux.Vars(r)["id"]
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println(errors.New("item not added"))
	// }
	// fmt.Println("string(body)", string(body))
	// // var itinerary *entity.Itinerary
	// resp, err := json.Marshal(body)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("resp", resp)
	// // fmt.Println("resp", resp)
	// // fmt.Println("itinerary", itinerary)
	res, err := h.service.GetItineraryById(id)
	fmt.Println("res", res)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("itinerary in handler", &itinerary)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Successfully added"))
	json.NewEncoder(w).Encode(res)
}

func (h *ItineraryHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	// id := mux.Vars(r)["itinerary_id"]
	// fmt.Println("id", id)
	// defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	fmt.Println("body", string(body))
	var item entity.ItineraryItem
	err = json.Unmarshal(body, &item)
	if err != nil {
		return
	}
	fmt.Println("unmarshal", item.ItineraryID)
	i, err := h.service.AddItem(item.ItineraryID, &item)
	fmt.Println(item)
	fmt.Println("i", i)
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
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}

func (h *ItineraryHandler) GetItemByItemId(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("id", id)
	var item *entity.ItineraryItem
	if id == "" {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
	fmt.Println("item", item)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func (h *ItineraryHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {}

func (h *ItineraryHandler) GetAllItineraries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler")
	data, err := h.service.GetAllItineraries()
	fmt.Println("data", data)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *ItineraryHandler) UpdateItinerary(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var updatedItinerary entity.Itinerary
	json.Unmarshal(body, &updatedItinerary)
	_, err = h.service.UpdateItineraryById(id, &updatedItinerary)
	if err != nil {
		return
	}

}

// func indexByID(items []entity.ItineraryItem, id string) int {
// 	for i := 0; i < len(items); i++ {
// 		if items[i].ItineraryID == id {
// 			return i
// 		}
// 	}
// 	return -1
// }
