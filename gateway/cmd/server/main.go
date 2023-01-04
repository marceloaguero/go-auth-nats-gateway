package main

import (
	"log"
	"os"

	"github.com/marceloaguero/go-auth-nats-gateway/gateway/pkg/delivery/router"
)

func main() {
	pathPrefix := os.Getenv("PATH_PREFIX")
	natsURLs := os.Getenv("NATS_URLS")

	_, err := router.NewRouter(pathPrefix, natsURLs)
	if err != nil {
		log.Panic(err)
	}
}
