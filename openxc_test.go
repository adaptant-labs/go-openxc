package openxc

import (
	"testing"
)

func TestOpenXC(t *testing.T) {
	var state VehicleState
	var msg VehicleMessage

	ds, err := OpenDataSource("trace", "trace-simple.json")
	if err != nil {
		t.Error("failed to open data source")
		t.FailNow()
	}

	defer ds.CloseDataSource()

	for {
		err := ds.ReadDataSource(&msg)
		if err != nil {
			break
		}

		err = VehicleMessageToState(&state, &msg)
		if err != nil {
			break
		}
	}

	// A simple sanity check to ensure that the vehicle state matches
	// matches the settings in the sample trace log
	if state.TransmissionGearPosition != "first" {
		t.Error("state does not match expectations")
		t.FailNow()
	}
}
