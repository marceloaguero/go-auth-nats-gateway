package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/marceloaguero/go-auth-nats-gateway/gateway/pkg/delivery/router"
)

func main() {
	pathPrefix := os.Getenv("PATH_PREFIX")
	natsURLs := os.Getenv("NATS_URLS")

	router, err := router.NewRouter(pathPrefix, natsURLs)
	if err != nil {
		log.Panic(err)
	}

	// Setup an interrupt handler to drain nats
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println()
	log.Printf("Draining...")
	router.Drain()
	log.Fatal("Exiting...")
}
