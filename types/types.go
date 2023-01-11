package types

// AirportCode stores the three letter IATA code of an airport
type AirportCode string

// Flight is a two element array of Source and Destination airports
type Flight [2]AirportCode

// FlightsData is an unordered list of flights in the itinerary of a person
type FlightsData []Flight

