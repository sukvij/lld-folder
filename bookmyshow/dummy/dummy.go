package dummy

import (
	"vijju/location"
	"vijju/movies"
	"vijju/shows"
	"vijju/theatre"
)

func DummyLocation() []*location.Location {
	locations := []*location.Location{}
	location := location.NewLocation("banglore")
	movies := movies.NewMovies("tiger 3")
	theatre := theatre.NewTheatre("galaxy theatre")
	shows := shows.NewShow("2 PM - 5 PM", 50)
	theatre.Shows = append(theatre.Shows, shows)
	movies.Theatres = append(movies.Theatres, theatre)
	location.Movies = append(location.Movies, movies)
	locations = append(locations, location)
	return locations
}
