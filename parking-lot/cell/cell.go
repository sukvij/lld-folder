package cell

import "vijju/vehicle"

type Cell struct {
	FloorNo     int
	CellNo      int
	Available   bool
	VehicleType *vehicle.VehicleType
	Vehicle     *vehicle.Vehicle
}

func NewCell(floor int, cellNo int) *Cell {
	return &Cell{FloorNo: floor, CellNo: cellNo, Available: true}
}
