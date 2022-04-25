package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// instatiate http server
// routes - mux

func main() {
	log.Printf("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
	router := mux.NewRouter()
	fmt.Println(router)
}
