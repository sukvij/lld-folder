package shows

import (
	"fmt"
	"vijju/seat"
)

type Show struct {
	Seats []*seat.Seat
	Time  string
}

func NewShow(time string, numerSeats int) *Show {
	seats := []*seat.Seat{}
	for i := 0; i < numerSeats; i++ {
		seat := seat.NewSeat(i)
		seats = append(seats, seat)
	}
	return &Show{Time: time, Seats: seats}
}

func (show *Show) PrintShow() {
	fmt.Println(show.Time)
	for i := 0; i < len(show.Seats); i++ {
		show.Seats[i].PrintSeat()
	}
}

func (show *Show) BookShow(seat *seat.Seat) {
	seat.BookSeat()
}
