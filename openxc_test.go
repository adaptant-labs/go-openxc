package openxc

import (
	"testing"
)

func TestOpenXC(t *testing.T) {
	var state VehicleState

	ds, err := OpenDataSource("trace", "trace-simple.json")
	if err != nil {
		t.Error("failed to open data source")
		t.FailNow()
	}

	defer ds.CloseDataSource()

	for {
		msg, err := ds.ReadDataSource()
		if err != nil {
			break
		}

		VehicleMessageToState(&state, msg)
	}

	// A simple sanity check to ensure that the vehicle state matches
	// matches the settings in the sample trace log
	if state.HeadlampStatus != true ||
		state.HighBeamStatus != false {
		t.Error("state does not match expectations")
		t.FailNow()
	}
}
