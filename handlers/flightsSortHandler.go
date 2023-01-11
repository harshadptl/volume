package handlers

import (
	"encoding/json"
	"github.com/harshadptl/volume_assignment/types"
	"log"
	"net/http"
)

type FlightsSortResponse struct {
	Success bool `json:"success"`
	Source string `json:"src"`
	Destination string `json:"dest"`
	Error string `json:"err"`
}
func FlightsSortHandler(res http.ResponseWriter, req *http.Request) {
	// parse the request
	fd, err := types.ParseFlightsData(req.Body)
	if err != nil {
		fsErrorResponse(res, err)
		return
	}

	log.Println(fd)
}

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
