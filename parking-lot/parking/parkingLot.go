package parking

import (
	"fmt"
	"vijju/cell"
	"vijju/token"
	"vijju/vehicle"
)

type ParkingLot struct {
	Floors             int
	ParkingInEachFloor int
	Lots               [][]*cell.Cell
}

func NewParkingLot(floor int, parkInEachFlorr int) *ParkingLot {
	parkingLot := &ParkingLot{Floors: floor, ParkingInEachFloor: parkInEachFlorr}
	lots := [][]*cell.Cell{}
	for i := 0; i < floor; i++ {
		mp := []*cell.Cell{}
		for j := 0; j < parkInEachFlorr; j++ {
			lot := cell.NewCell(i, j)
			if j%3 == 0 {
				lot.VehicleType = &vehicle.TypeOfWehicle.TwoWheeler
			} else if j%3 == 1 {
				lot.VehicleType = &vehicle.TypeOfWehicle.ThreeWheeler
			} else {
				lot.VehicleType = &vehicle.TypeOfWehicle.FourWheeler
			}
			mp = append(mp, lot)
		}
		lots = append(lots, mp)
	}
	parkingLot.Lots = lots
	return parkingLot
}

func (parkingLot *ParkingLot) ParkVehicle(vehicle *vehicle.Vehicle) (*token.ParkingToken, bool) {

	for i := 0; i < parkingLot.Floors; i++ {
		for j := 0; j < parkingLot.ParkingInEachFloor; j++ {
			lot := parkingLot.Lots[i][j]
			if vehicle.Type == *lot.VehicleType && lot.Available {
				lot.Available = false
				token := token.NewParkingToken(i, j)
				fmt.Println("vehicle parked at place ", token.FloorNo, " ", token.CellNo)
				return token, false
			}
		}
	}
	fmt.Println("all palces are booked")
	return nil, true
}
