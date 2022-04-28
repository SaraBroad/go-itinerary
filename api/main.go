package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SaraBroad/go-itinerary/api/repository"
	"github.com/gorilla/mux"
)

// instatiate http server
// routes - mux

func main() {
	port := "8080"
	log.Printf("Server started at port %v", port)
	http.ListenAndServe(":"+port, nil)
	_ = repository.InitDatabase(&repository.Auth{})
	_ = repository.NewItinerary()
	router := mux.NewRouter()
	fmt.Println(router)
	// router.HandleFunc("/", i.CreateNewItem(&models.Item{})).Methods(http.MethodGet)
}

// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter text: ")
// text, _ := reader.ReadString('\n')
// fmt.Println(text)
