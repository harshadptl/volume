package types

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
)

// Pattern to verify Airport Code
var pattern = regexp.MustCompile("^\"[A-Z]{3}\"$")

//
// AirportCode stores the three letter IATA code of an airport
//
type AirportCode string

// While unmarshalling Airport code check if it is 3 character long string
// with surrounding double quotes. It returns an error if the format is wrong.
func (d *AirportCode) UnmarshalJSON(data []byte) error {
	if !pattern.Match(data) {
		return ErrInvalidAirportCode
	}
	*d = AirportCode(data[1:4])
	return nil
}


//
// Flight is a two element array of Source and Destination airports
//
type Flight [2]AirportCode

func (fl *Flight) Data() map[string]string {
	return map[string]string{
		"src": string(fl[0]),
		"dest": string(fl[1]),
	}
}

//
// FlightsData is an unordered list of flights in the itinerary of a person
//
type FlightsData []Flight

// ParseFlightsData parses the request body for FlightsData object. It returns
// an error if the format is wrong.
func ParseFlightsData(r io.Reader) (*FlightsData, error) {
	var f FlightsData
	err := json.NewDecoder(r).Decode(&f)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing FlightsData: %s", err)
	}
	return &f, nil
}

