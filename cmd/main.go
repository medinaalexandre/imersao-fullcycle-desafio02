package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpHandler "github.com/medinaalexandre/imersao-fullcycle-desafio02/internal/infra/http"
	"github.com/medinaalexandre/imersao-fullcycle-desafio02/internal/infra/repository"
	"github.com/medinaalexandre/imersao-fullcycle-desafio02/internal/usecase"
)

func main() {
	repo, err := repository.NewJsonEventRepository("data.json")

	if err != nil {
		print(err.Error())
		return
	}

	eventsHandler := httpHandler.NewEventsHandler(
		usecase.NewListEventsUseCase(repo),
		usecase.NewGetEventUseCase(repo),
		usecase.NewListSpotsUseCase(repo),
		usecase.NewReserveSpotUseCase(repo),
	)

	r := http.NewServeMux()
	r.HandleFunc("/events", eventsHandler.ListEvents)
	r.HandleFunc("/events/{eventID}", eventsHandler.GetEvent)
	r.HandleFunc("/events/{eventID}/spots", eventsHandler.ListSpotEvent)
	r.HandleFunc("POST /events/{eventID}/reserve", eventsHandler.ReserveSpot)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Canal para escutar sinais do sistema operacional
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// Recebido sinal de interrupção, iniciando o graceful shutdown
		log.Println("Recebido sinal de interrupção, iniciando o graceful shutdown...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Erro no graceful shutdown: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Iniciando o servidor HTTP
	log.Println("Servidor HTTP rodando na porta 8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor HTTP: %v\n", err)
	}

	<-idleConnsClosed
	log.Println("Servidor HTTP finalizado")
}
