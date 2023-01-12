package types

import "errors"

// ErrInvalidAirportCode is used to raise an error for bad Airport code format
var ErrInvalidAirportCode = errors.New("Invalid Airport Code")

// ErrInvalidFlightList is raised when the flight list doesn't have a valid sequence
var ErrInvalidFlightList = errors.New("Invalid Flight List")
