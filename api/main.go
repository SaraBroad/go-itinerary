package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SaraBroad/go-itinerary/api/domain/repository"
	"github.com/SaraBroad/go-itinerary/api/handler"
	"github.com/SaraBroad/go-itinerary/api/service"
	"github.com/gorilla/mux"
)

func main() {
	db := repository.ConnectDatabase()
	// s := seed.Database{}
	// _, err := s.Seed()
	// if err != nil {
	// 	fmt.Println("seed error")
	// }
	ir := repository.NewItinerary(db)
	fmt.Println("ir", ir)
	is := service.NewItineraryService(ir)
	ih := handler.NewItemHandler(*is)
	// fmt.Println("ih", ih.GetItinerary)
	router := mux.NewRouter()
	// router.HandleFunc("/items", ih.AddItem).Methods("POST")
	// router.HandleFunc("/items/{id}", ih.RemoveItem).Methods("DELETE")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/itinerary", ih.AddItinerary).Methods("POST")
	router.HandleFunc("/itinerary/{id}", ih.GetItinerary).Methods("GET")
	http.Handle("/", router)

	log.Println("http server runs on :8080")
	http.ListenAndServe(":8080", router)
}

// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter text: ")
// text, _ := reader.ReadString('\n')
// fmt.Println(text)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
