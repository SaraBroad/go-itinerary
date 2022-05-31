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
	fmt.Println("r", r)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("body", body)
	var item models.Item
	json.Unmarshal(body, &item)

	err = h.service.AddItem(&item)
	fmt.Println(item)
	if err != nil {
		fmt.Println(errors.New("item not added"))
	}
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Successfully Updated"))
}
