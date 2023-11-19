package main

import (
	"vijju/system"
	"vijju/vehicle"
)

func main() {
	system := system.NewSystem()
	vehicle1 := vehicle.NewVehicle("bmw", vehicle.TypeOfWehicle.FourWheeler)
	vehicle2 := vehicle.NewVehicle("bmw", vehicle.TypeOfWehicle.FourWheeler)
	vehicle3 := vehicle.NewVehicle("bmw", vehicle.TypeOfWehicle.TwoWheeler)
	system.EntryVehicle(vehicle1)
	system.EntryVehicle(vehicle2)
	system.EntryVehicle(vehicle3)
}
