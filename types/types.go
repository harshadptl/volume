package types

import (
	"encoding/json"
	"fmt"
	"io"
)


// AirportCode stores the three letter IATA code of an airport
type AirportCode string

// Flight is a two element array of Source and Destination airports
type Flight [2]AirportCode

// FlightsData is an unordered list of flights in the itinerary of a person
type FlightsData []Flight

func ParseFlightsData(r io.Reader) (*FlightsData, error) {
	var f FlightsData
	err := json.NewDecoder(r).Decode(&f)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing FlightsData: %s", err)
	}
	return &f, nil
}

// While unmarshalling Airport code check if it is 3 character long surrounded by quotes
func (d *AirportCode) UnmarshalJSON(data []byte) error {
	fmt.Println(data)
	if len(data) != 5 || data[0]!=34 || data[4]!=34 {
		return ErrInvalidAirportCode
	}
	*d = AirportCode(data[1:4])
	return nil
}