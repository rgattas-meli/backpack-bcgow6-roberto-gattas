package service

import "errors"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, v := range b.Tickets {
		if v.Id == t.Id {
			return t, errors.New("Id ya existe")
		}

		b.Tickets = append(b.Tickets, t)
	}
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, v := range b.Tickets {
		if v.Id == id {
			return v, nil
		}
	}
	return Ticket{}, errors.New("id incorrecto")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, v := range b.Tickets {
		if v.Id == id {
			b.Tickets[i] = t
			return t, nil
		}
	}
	return Ticket{}, errors.New("id incorrecto")
}

func (b *bookings) Delete(id int) (int, error) {
	for i, v := range b.Tickets {
		if v.Id == id {
			b.Tickets[i] = b.Tickets[len(b.Tickets)-1]
			b.Tickets = b.Tickets[:len(b.Tickets)-1]
			return id, nil

		}
	}
	return 0, errors.New("id incorrecto")
}
