package movies

import (
	"fmt"
	"vijju/theatre"
)

type Movies struct {
	Name     string
	Theatres []*theatre.Theatre
}

func NewMovies(name string) *Movies {
	theatres := []*theatre.Theatre{}
	return &Movies{Name: name, Theatres: theatres}
}

func (movie *Movies) PrintMovie() {
	fmt.Println(movie.Name)
	// for i := 0; i < len(movie.Theatres); i++ {
	// 	movie.Theatres[i].PrintTheatre()
	// }
}
