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
	port := "8080"
	log.Printf("Server started at port %v", port)
	http.ListenAndServe(":"+port, nil)
	router := mux.NewRouter()
	fmt.Println(router)
}
