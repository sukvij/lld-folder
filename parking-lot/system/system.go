package system

import (
	"vijju/entry"
	"vijju/parking"
	"vijju/token"
	"vijju/vehicle"
)

type System struct {
	ParkingLot *parking.ParkingLot
	Gates      []*entry.EntryGate
}

func NewSystem() *System {
	system := &System{}
	system.ParkingLot = parking.NewParkingLot(10, 10)
	system.Gates = entry.EntryGates(5)
	return system
}

func (system *System) EntryVehicle(vehicle *vehicle.Vehicle) (*token.ParkingToken, bool) {
	parkingLot := system.ParkingLot
	token, err := parkingLot.ParkVehicle(vehicle)
	if !err {
		return token, false
	}
	return nil, true
}
