package seat

import "fmt"

type Seat struct {
	Id        int
	Available bool
}

func NewSeat(id int) *Seat {
	return &Seat{Id: id, Available: true}
}

func (seat *Seat) BookSeat() {
	fmt.Println("user locked this seat")
	seat.Available = false
	fmt.Println("user booked seat")
}

func (seat *Seat) PrintSeat() {
	fmt.Println(seat)
}
