package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/BetuelSA/go-helpers/password"
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/delivery"
	repo "github.com/marceloaguero/go-auth-nats-gateway/users/pkg/repository"
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/user"
)

func main() {
	dbDsn := os.Getenv("DB_DSN")
	dbName := os.Getenv("DB_NAME")
	natsURLs := os.Getenv("NATS_URLS")
	subjPrefix := os.Getenv("SUBJ_PREFIX")
	queue := os.Getenv("QUEUE")

	repository, err := repo.NewRepo(dbDsn, dbName)
	if err != nil {
		log.Panic(err)
	}

	// Usecase uses a password service
	pass := password.NewService()
	usecase := user.NewUsecase(repository, pass)

	delivery, err := delivery.NewDelivery(usecase, natsURLs, subjPrefix, queue)
	if err != nil {
		log.Panic(err)
	}

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("Draining...")
	delivery.Drain()
	log.Fatalf("Exiting")
}
