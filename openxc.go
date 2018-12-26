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
	"strconv"
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
	Value string
	Event bool
}

// VehicleMessageToState updates the VehicleState information with the contents
// of a single VehicleMessage.
func VehicleMessageToState(state *VehicleState, msg *VehicleMessage) error {
	var err error = nil

	switch msg.Name {
	case "headlamp_status":
		state.HeadlampStatus, err = strconv.ParseBool(msg.Value)
	case "high_beam_status":
		state.HighBeamStatus, err = strconv.ParseBool(msg.Value)
	case "brake_pedal_status":
		state.BrakePedalStatus, err = strconv.ParseBool(msg.Value)
	case "windshield_wiper_status":
		state.WindshieldWiperStatus, err = strconv.ParseBool(msg.Value)
	case "parking_brake_status":
		state.ParkingBrakeStatus, err = strconv.ParseBool(msg.Value)
	case "door_status":
		state.DoorStatus = msg.Value
	case "ignition_status":
		state.IgnitionStatus = msg.Value
	case "gear_lever_position":
		state.GearLeverPosition = msg.Value
	case "transmission_gear_position":
		state.TransmissionGearPosition = msg.Value
	case "fuel_level":
		state.FuelLevel, err = strconv.ParseFloat(msg.Value, 64)
	case "latitude":
		state.Latitude, err = strconv.ParseFloat(msg.Value, 64)
	case "longitude":
		state.Longitude, err = strconv.ParseFloat(msg.Value, 64)
	case "accelerator_pedal_position":
		state.AcceleratorPedalPosition, err = strconv.ParseFloat(msg.Value, 64)
	case "engine_speed":
		state.EngineSpeed, err = strconv.ParseFloat(msg.Value, 64)
	case "vehicle_speed":
		state.VehicleSpeed, err = strconv.ParseFloat(msg.Value, 64)
	case "fuel_consumed_since_restart":
		state.FuelConsumedSinceRestart, err = strconv.ParseFloat(msg.Value, 64)
	case "odometer":
		state.Odometer, err = strconv.ParseFloat(msg.Value, 64)
	case "steering_wheel_angle":
		state.SteeringWheelAngle, err = strconv.ParseFloat(msg.Value, 64)
	case "torque_at_transmission":
		state.TorqueAtTransmission, err = strconv.ParseFloat(msg.Value, 64)
	default:
		fmt.Println("unhandled message for type", msg.Name)
	}

	return err
}
