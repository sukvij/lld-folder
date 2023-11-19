package theatre

import (
	"fmt"
	"vijju/shows"
)

type Theatre struct {
	Name  string
	Shows []*shows.Show
}

func NewTheatre(name string) *Theatre {
	shows := []*shows.Show{}
	return &Theatre{Name: name, Shows: shows}
}

func (theatre *Theatre) PrintTheatre() {
	fmt.Println(theatre.Name)
	// for i := 0; i < len(theatre.Shows); i++ {
	// 	theatre.Shows[i].PrintShow()
	// }
}
