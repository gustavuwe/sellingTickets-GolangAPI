package domain

import (
	"errors"
	"time"
)

var (
	ErrEventNameRequired = errors.New("event name is required")
	ErrEventDateFuture   = errors.New("event date must be in the future")
	ErrEventCapacityZero = errors.New("event capacity must be greater than zero")
	ErrEventPriceZero    = errors.New("event price must be greater than zero")
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "L10"
	Rating12    Rating = "L12"
	Rating14    Rating = "L14"
	Rating16    Rating = "L16"
	Rating18    Rating = "L18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    int
	Spots        []Spot
	Tickets      []Ticket
}

func (e Event) Validate() error { // e *Event the same thing of: (e): Event on typescript, so im passing the type Event to my parameter
	if e.Name == "" {
		return ErrEventNameRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrEventDateFuture
	}

	if e.Capacity <= 0 {
		return ErrEventCapacityZero
	}

	if e.Price <= 0 {
		return ErrEventPriceZero
	}

	return nil // inexistent value, or no errors found,
}

func (e *Event) AddSpot(name string) (*Spot, error) { // after the name of function is what the function gonna return, so in this case gona return a spot or an error.
	spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}

	e.Spots = append(e.Spots, *spot)
	return spot, nil
}
