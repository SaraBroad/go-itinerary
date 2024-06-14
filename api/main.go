package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"net/http/httputil"

	"github.com/SaraBroad/go-itinerary/api/domain/repository"
	"github.com/SaraBroad/go-itinerary/api/handler"
	"github.com/SaraBroad/go-itinerary/api/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("here")
	db := repository.ConnectDatabase()
	// s := seed.Database{}
	// _, err := s.Seed()
	// if err != nil {
	// 	fmt.Println("seed error")
	// }
	ir := repository.NewItinerary(db)
	is := service.NewItineraryService(ir)
	ih := handler.NewItemHandler(*is)
	router := mux.NewRouter()
	// route := http.NewServeMux()
	// fmt.Println(router)
	// router.Use("/")
	// route.HandleFunc("/itinerary", handleRequest)
	// router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/add-item", ih.AddItem).Methods("POST")
	// // router.HandleFunc("/items/{id}", ih.RemoveItem).Methods("DELETE")
	// router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/itinerary", ih.AddItinerary).Methods("POST")
	router.HandleFunc("/itinerary/{id}", ih.GetItinerary).Methods("GET")
	router.HandleFunc("/itineraries", ih.GetAllItineraries).Methods("GET")
	http.Handle("/", handlers.CORS()(router))

	log.Println("http server runs on :8080")
	err := http.ListenAndServe(":8080", handlers.CORS()(router))
	if err != nil {
		log.Fatal(err)
	}
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

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, Accept")
		} else {
			h.ServeHTTP(w, r)
		}
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		w.WriteHeader(http.StatusOK)
		return
	}
	targetUrl, _ := url.Parse("http://localhost:3000")

	proxy := httputil.NewSingleHostReverseProxy(targetUrl)
	proxy.ServeHTTP(w, r)
}
