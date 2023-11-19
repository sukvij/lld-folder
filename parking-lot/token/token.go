package token

type ParkingToken struct {
	FloorNo int
	CellNo  int
}

func NewParkingToken(floor int, cellNo int) *ParkingToken {
	return &ParkingToken{FloorNo: floor, CellNo: cellNo}
}
