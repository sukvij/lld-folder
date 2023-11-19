package system

import (
	"vijju/dummy"
	"vijju/location"
)

type System struct {
	Location []*location.Location
}

func NewSystem() *System {
	location := dummy.DummyLocation()
	return &System{Location: location}
}
