package types

import "testing"



func TestAirportCode_UnmarshalJSON(t *testing.T) {
	code := AirportCode("")
	err := code.UnmarshalJSON([]byte(`"JFK"`))
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}

	badCodes := []string{`"JFK_"`, `JFK`, `asd`, `123`, `ErE`, `"asd"`, `"123"`, `"ErE"`, `" JFK"`}
	for _, c := range badCodes {
		err = code.UnmarshalJSON([]byte(c))
		if err != ErrInvalidAirportCode {
			t.Log("failed to verify airport code", err)
			t.Fail()
		}
	}
}

func TestFlight_Data(t *testing.T) {
	f := Flight{AirportCode("ABC"), AirportCode("XYZ")}
	if f.Data()["src"] != "ABC" {
		t.Log("bad parsing")
		t.Fail()
	}
	if f.Data()["dest"] != "XYZ" {
		t.Log("bad parsing")
		t.Fail()
	}
}
