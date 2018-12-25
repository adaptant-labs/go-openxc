// Copyright 2018 Adaptant Labs
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// The openxc package provides a simple API for working with data from the
// OpenXC platform.
package openxc

import (
	"fmt"
)

// VehicleState contains basic state information about a given vehicle.
type VehicleState struct {
	HeadlampStatus			bool
	HighBeamStatus			bool
	WindshieldWiperStatus		bool
	BrakePedalStatus		bool

	DoorStatus			string
	IgnitionStatus			string
	TurnSignalStatus		string

	GearLevelPosition		string
	TransmissionGearPosition	string
}

// VehicleMessage is the top-level wrapper for the OpenXC messaging format. Two
// types of messages, according to their JSON encoding, are dealt with by the
// specification:
//
// Simple vehicle messages:
//	{"name": "headlamp_status", "value": false}
//
// Extended vehicle messages:
//	{"name": "headlamp_status", "value": false, "event": false}
type VehicleMessage struct {
	Name	string		`json:"name",required`
	Value	interface{}	`json:"value",required`
	Event	bool		`json:"event",omitempty"`
}

// VehicleMessageToState updates the VehicleState information with the contents
// of a single VehicleMessage.
func VehicleMessageToState(state *VehicleState, msg VehicleMessage) {
	switch msg.Name {
	case "headlamp_status":
		state.HeadlampStatus = msg.Value.(bool);
	case "high_beam_status":
		state.HighBeamStatus = msg.Value.(bool);
	case "brake_pedal_status":
		state.BrakePedalStatus = msg.Value.(bool);
	case "windshield_wiper_status":
		state.WindshieldWiperStatus = msg.Value.(bool);
	case "door_status":
		state.DoorStatus = msg.Value.(string);
	case "ignition_status":
		state.IgnitionStatus = msg.Value.(string);
	default:
		fmt.Println("unhandled message for type", msg.Name)
	}
}
