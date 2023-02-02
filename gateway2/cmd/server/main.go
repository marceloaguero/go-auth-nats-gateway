package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/marceloaguero/go-auth-nats-gateway/gateway2/pkg/delivery/router"
	"github.com/marceloaguero/go-auth-nats-gateway/gateway2/pkg/delivery/users"
	"github.com/nats-io/nats.go"
)

func main() {
	pathPrefix := os.Getenv("PATH_PREFIX")
	natsURLs := os.Getenv("NATS_URLS")
	usersSubjPrefix := os.Getenv("USERS_SUBJ_PREFIX")
	usersQueue := os.Getenv("USERS_QUEUE")

	// Connect to NATS server
	nc, err := nats.Connect(natsURLs)
	if err != nil {
		log.Panic(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Panic(err)
	}
	defer ec.Close()

	usersDelivery := users.NewDelivery(ec, usersSubjPrefix, usersQueue)

	_, err = router.NewRouter(usersDelivery, pathPrefix)
	if err != nil {
		log.Panic(err)
	}

	// Setup an interrupt handler to drain nats
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	nc.Drain()
	log.Fatal("Exiting...")
}
