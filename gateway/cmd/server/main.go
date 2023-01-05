package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/marceloaguero/go-auth-nats-gateway/gateway/pkg/delivery/router"
	"github.com/marceloaguero/go-auth-nats-gateway/gateway/pkg/delivery/users"
	"github.com/nats-io/nats.go"
)

func main() {
	pathPrefix := os.Getenv("PATH_PREFIX")
	natsURLs := os.Getenv("NATS_URLS")

	// Connect to NATS server
	nc, err := nats.Connect(natsURLs)
	if err != nil {
		log.Panic(err)
	}
	defer nc.Close()

	usersDelivery := users.NewDelivery(nc, "USERS", "users")

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
