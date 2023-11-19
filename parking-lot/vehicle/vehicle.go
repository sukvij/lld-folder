package vehicle

type VehicleType string

var TypeOfWehicle = struct {
	TwoWheeler   VehicleType
	ThreeWheeler VehicleType
	FourWheeler  VehicleType
}{
	TwoWheeler:   "two wheeler",
	ThreeWheeler: "three wheeler",
	FourWheeler:  "four wheeler",
}

type Vehicle struct {
	Name string
	Type VehicleType
}

func NewVehicle(name string, typeVehicle VehicleType) *Vehicle {
	return &Vehicle{Name: name, Type: typeVehicle}
}
