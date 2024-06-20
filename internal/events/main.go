package main

import (
	"database/sql"
	"net/http"

	"github.com/devfullcycle/imersao18/golang/internal/events/infra/repository"
	"github.com/devfullcycle/imersao18/golang/internal/events/infra/service"
	"github.com/devfullcycle/imersao18/golang/internal/events/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test_user:test_password@tcp(localhost:3306)")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventRepo, err := repository.NewMysqlEventRepository(db)
	if err != nil {
		panic(err)
	}

	// partnerBaseURLs := map[int]string{
	// 	1: "http://localhost:9080/api1",
	// 	2: "http://localhost:9080/api2",
	// }

	// partnerFactory := service.NewPartnerFactory(partnerBaseURLs)

	listEventsUseCase := usecase.NewListEventUseCase(eventRepo)
	getEventUseCase := usecase.NewGetEventUseCase(eventRepo)
	listSpotsUseCase := usecase.NewListEventUseCase(eventRepo)
	// BuyTicketsUseCase := usecase.NewBuyTicketsUseCase(eventRepo, partnerFactory)

	eventsHandler := httpHandler.NewEventHandler(
		listEventsUseCase,
		listSpotsUseCase,
		getEventUseCase,
		// BuyTicketsUseCase,
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpots)
	r.HandleFunc("POST /checkout", eventsHandler.BuyTickets)
}
