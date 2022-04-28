package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SaraBroad/go-itinerary/api/repository"
)

// instatiate http server
// routes - mux

var port = "8080"

func main() {

	_ = repository.InitDatabase(&repository.Auth{
		Username: os.Getenv("USERNAME"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		Port:     os.Getenv("PORT"),
		URI:      os.Getenv("URL"),
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
