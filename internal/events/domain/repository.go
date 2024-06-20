package domain

// interface of functions work like this: functionA() (argB, error) -> so the functionA() when executed will return the argB if success or an error.

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(eventID string) (*Event, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	FindSpotByName(eventID, spotName string) (*Spot, error)
	// CreateEvent(event *Event) error
	CreateSpot(spot *Spot) error
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, TicketID string) error
}
