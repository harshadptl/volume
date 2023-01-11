package main

import (
	"github.com/harshadptl/volume_assignment/handlers"
	"log"
	"net/http"
)

// PORT on which the http server listens
const PORT = ":8080"

func main() {
	//Create the default mux
	mux := http.NewServeMux()

	//Handling the /v1/flights/sort. The handler is a function here
	mux.HandleFunc("/v1/flights/sort", handlers.FlightsSortHandler)

	//Create the server.
	s := &http.Server{
		Addr:    PORT,
		Handler: mux,
	}
	err := s.ListenAndServe()
	log.Println("ListenAndServe error: ", err)
}
