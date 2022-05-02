package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SaraBroad/go-itinerary/api/repository"
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

	_ = repository.InitDatabase(&repository.Auth{
		Username: os.Getenv("USERNAME"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		Port:     os.Getenv("PORT"),
		SSL:      os.Getenv("SSL"),
		Timezone: os.Getenv("TIMEZONE"),
	})
	_ = repository.NewItinerary()
	// router := mux.NewRouter()
	log.Printf("Server started at port %v", port)
	http.ListenAndServe(":"+port, nil)
	// router.HandleFunc("/", i.CreateNewItem(&models.Item{})).Methods(http.MethodGet)
}

// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter text: ")
// text, _ := reader.ReadString('\n')
// fmt.Println(text)
