package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketType string

var (
	ErrTicketPriceZero = errors.New("ticket price must be greater than zero")
)

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TicketTypeFull TicketType = "full" // Full-price ticket
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot // *: when changed value of spot, gonna change the value in the whole program
	TicketType TicketType
	Price      float64
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !isValidTicketType(ticketType) {
		return nil, errors.New("invalid ticket type")
	}
	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}

	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}

	return nil
}
