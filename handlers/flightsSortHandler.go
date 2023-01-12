package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/harshadptl/volume_assignment/flights"
	"github.com/harshadptl/volume_assignment/types"
)

// FlightsSortResponse is used to generate the JSON response for the FLights Sort API
type FlightsSortResponse struct {
	Success bool `json:"success"`
	Source string `json:"src"`
	Destination string `json:"dest"`
	Error string `json:"err"`
}

// FlightsSortHandler implements the logic for handling Flights Sort request
func FlightsSortHandler(res http.ResponseWriter, req *http.Request) {

	// Catch panic and raise 500 status response
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
			res.WriteHeader(http.StatusInternalServerError)
		}
	}()

	// parse the request
	fd, err := types.ParseFlightsData(req.Body)
	if err != nil {
		fsErrorResponse(res, err)
		return
	}

	log.Println("Unsorted Flights list request: ", fd)

	// Sort the flight data with helper, return error if there is one
	list, err := flights.FlightSort(fd)
	if err != nil {
		fsErrorResponse(res, err)
		return
	}

	// Return the non nil list in the collection as the final sorted flight list
	fsSuccessResponse(
		res,
		list.Head.Data["src"],
		list.Tail.Data["dest"],
	)
}

// fsErrorResponse writes the Bad Request error response to the writer
func fsErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(
		FlightsSortResponse{
			Success:     false,
			Source:      "",
			Destination: "",
			Error:       err.Error(),
		})
}

// fsSuccessResponse writes the Success response of API to the response writer
func fsSuccessResponse(w http.ResponseWriter, src, dest string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		FlightsSortResponse{
			Success:     true,
			Source:      src,
			Destination: dest,
			Error:       "",
		})
}