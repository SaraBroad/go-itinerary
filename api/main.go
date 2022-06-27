package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SaraBroad/go-itinerary/api/handler"
	"github.com/SaraBroad/go-itinerary/api/repository"
	"github.com/SaraBroad/go-itinerary/api/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// instatiate http server
// routes - mux

var port = "8080"

func main() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := repository.InitDatabase(&repository.Auth{
		Username: os.Getenv("USERNAME"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		Port:     os.Getenv("PORT"),
		SSL:      os.Getenv("SSL"),
		Timezone: os.Getenv("TIMEZONE"),
	})

	ir := repository.NewItinerary(db)
	is := service.NewItemService(ir)
	ih := handler.NewItemHandler(*is)

	router := mux.NewRouter()
	router.HandleFunc("/items", ih.AddItem).Methods("POST")
	router.HandleFunc("/items/{id}", ih.RemoveItem).Methods("DELETE")
	http.Handle("/", router)

	log.Println("http server runs on :8080")
	http.ListenAndServe(":8080", router)

	log.Printf("Server started at port %v", port)
}

// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter text: ")
// text, _ := reader.ReadString('\n')
// fmt.Println(text)
