package flights

import (
	"fmt"
	"testing"

	"github.com/harshadptl/volume_assignment/types"
)

var testData = map[string]types.FlightsData{
	"ABC-XYZ": types.FlightsData{types.Flight{"ABC","XYZ"}},
	"LON-SFO": types.FlightsData{types.Flight{"JFK","SFO"}, types.Flight{"LON", "JFK"}},
	"BOM-SFO": types.FlightsData{types.Flight{"JFK","SFO"}, types.Flight{"LON", "JFK"}, types.Flight{"BOM", "LON"}},
	"SIN-YVR": types.FlightsData{types.Flight{"JFK","SFO"}, types.Flight{"LON", "JFK"}, types.Flight{"SIN", "LON"}, types.Flight{"SFO", "YVR"}},
}

var failingtestData = map[string]types.FlightsData{
	"LON-SFO": types.FlightsData{types.Flight{"JFK","SFO"}, types.Flight{"JFK", "SIN"}},
	"BOM-SFO": types.FlightsData{types.Flight{"JFK","SFO"}, types.Flight{"LON", "ADB"}, types.Flight{"BOM", "LON"}},
	"SIN-YVR": types.FlightsData{types.Flight{"JFK","SFO"}, types.Flight{"SIN", "LON"}, types.Flight{"SFO", "YVR"}},
}

func TestFlightSort(t *testing.T) {
	for k, v := range testData {

		l, err := FlightSort(&v)
		if err != nil {
			t.Log("Flight Sort error: ", err)
			t.Fail()
		}
		s := fmt.Sprintf("%s-%s", l.Head.Data["src"], l.Tail.Data["dest"])
		if s != k {
			t.Log("incorrect SRC DEST:", k, s)
			t.Fail()
		}
	}

	for _, v := range failingtestData {

		_, err := FlightSort(&v)
		if err != types.ErrInvalidFlightList {
			t.Log("something wrong:  ", err)
			t.Fail()
		}
	}
}

