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

// Package openxc provides a simple API for working with data from the OpenXC
// platform.
package openxc

import (
	"fmt"
)

// VehicleState contains basic state information about a given vehicle.
type VehicleState struct {
	HeadlampStatus        bool
	HighBeamStatus        bool
	WindshieldWiperStatus bool
	BrakePedalStatus      bool
	ParkingBrakeStatus    bool

	DoorStatus       string
	IgnitionStatus   string
	TurnSignalStatus string

	GearLeverPosition        string
	TransmissionGearPosition string

	FuelLevel                float64
	Latitude                 float64
	Longitude                float64
	AcceleratorPedalPosition float64
	EngineSpeed              float64
	VehicleSpeed             float64
	FuelConsumedSinceRestart float64
	Odometer                 float64
	SteeringWheelAngle       float64
	TorqueAtTransmission     float64
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
	Name  string
	Value interface{}
	Event bool
}

// VehicleMessageToState updates the VehicleState information with the contents
// of a single VehicleMessage.
func VehicleMessageToState(state *VehicleState, msg VehicleMessage) {
	switch msg.Name {
	case "headlamp_status":
		state.HeadlampStatus = msg.Value.(bool)
	case "high_beam_status":
		state.HighBeamStatus = msg.Value.(bool)
	case "brake_pedal_status":
		state.BrakePedalStatus = msg.Value.(bool)
	case "windshield_wiper_status":
		state.WindshieldWiperStatus = msg.Value.(bool)
	case "parking_brake_status":
		state.ParkingBrakeStatus = msg.Value.(bool)
	case "door_status":
		state.DoorStatus = msg.Value.(string)
	case "ignition_status":
		state.IgnitionStatus = msg.Value.(string)
	case "gear_lever_position":
		state.GearLeverPosition = msg.Value.(string)
	case "transmission_gear_position":
		state.TransmissionGearPosition = msg.Value.(string)
	case "fuel_level":
		state.FuelLevel = msg.Value.(float64)
	case "latitude":
		state.Latitude = msg.Value.(float64)
	case "longitude":
		state.Longitude = msg.Value.(float64)
	case "accelerator_pedal_position":
		state.AcceleratorPedalPosition = msg.Value.(float64)
	case "engine_speed":
		state.EngineSpeed = msg.Value.(float64)
	case "vehicle_speed":
		state.VehicleSpeed = msg.Value.(float64)
	case "fuel_consumed_since_restart":
		state.FuelConsumedSinceRestart = msg.Value.(float64)
	case "odometer":
		state.Odometer = msg.Value.(float64)
	case "steering_wheel_angle":
		state.SteeringWheelAngle = msg.Value.(float64)
	case "torque_at_transmission":
		state.TorqueAtTransmission = msg.Value.(float64)
	default:
		fmt.Println("unhandled message for type", msg.Name)
	}
}
