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

	// i := &models.Item{
	// 	Name:  "Hyatt",
	// 	Price: 199.99,
	// 	Category: models.Category{
	// 		Name: "Accomodations",
	// 	},
	// 	DayNumber: models.DayNumber{
	// 		DayNum: 1,
	// 	},
	// }

	// db := repository.Database{}
	// _, _ = ir.CreateNewItem(i)
	router := mux.NewRouter()
	router.HandleFunc("/items", ih.AddItem)

	log.Printf("Server started at port %v", port)
	http.ListenAndServe(":"+port, nil)
	// router.HandleFunc("/", i.CreateNewItem(&models.Item{})).Methods(http.MethodGet)
}

// reader := bufio.NewReader(os.Stdin)
// fmt.Print("Enter text: ")
// text, _ := reader.ReadString('\n')
// fmt.Println(text)
