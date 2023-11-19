package system

import "fmt"

func (system *System) BookShow() {
	fmt.Println("all locations")
	for i := 0; i < len(system.Location); i++ {
		fmt.Println(system.Location[i].Name)
	}

	fmt.Println()

	var location int
	fmt.Println("enter location")
	fmt.Scanln(&location)
	fmt.Println("all movies in location ", location)
	movies := system.Location[location].Movies
	for i := 0; i < len(movies); i++ {
		fmt.Println(i, " - ", movies[i].Name)
	}

	fmt.Println()
	fmt.Println("select movie")
	var movieId int
	fmt.Scanln(&movieId)
	theatres := movies[movieId].Theatres
	fmt.Println("all theatres")
	for i := 0; i < len(theatres); i++ {
		fmt.Println(i, " - ", theatres[i].Name)
	}

	fmt.Println()
	fmt.Println("Select theatreid")
	var theatreId int
	fmt.Scanln(&theatreId)
	shows := theatres[theatreId].Shows
	fmt.Println("all shows are given")
	for i := 0; i < len(shows); i++ {
		fmt.Println(i, " - ", shows[i].Time)
	}

	fmt.Println()
	fmt.Println("Select time for show")
	var showId int
	fmt.Scanln(&showId)
	shows[showId].PrintShow()
}
