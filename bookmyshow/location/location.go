package location

import (
	"fmt"
	"vijju/movies"
)

type Location struct {
	Name   string
	Movies []*movies.Movies
}

func NewLocation(name string) *Location {
	movies := []*movies.Movies{}
	return &Location{Name: name, Movies: movies}
}

func (location *Location) PrintAllLocation() {
	fmt.Println(location.Name)
	// for i := 0; i < len(location.Movies); i++ {
	// 	location.Movies[i].PrintMovie()
	// }
}
