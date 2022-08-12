package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/BetuelSA/go-helpers/password"
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/delivery"
	repo "github.com/marceloaguero/go-auth-nats-gateway/users/pkg/repository"
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/user"
	"github.com/nats-io/nats.go"
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

	// Connectto NATS
	nc, err := nats.Connect(natsURLs)
	if err != nil {
		log.Panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	err = delivery.Subscribe(usecase, ec, subjPrefix, queue)
	if err != nil {
		log.Panic(err)
	}

	// Setup the interrupt handler to drain so we don't miss
	// requests when scaling down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatalf("Exiting")
}
