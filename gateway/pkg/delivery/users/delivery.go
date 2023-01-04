package users

import (
	"time"

	"github.com/nats-io/nats.go"
)

const (
	timeout = time.Second * 2
)

type Delivery interface {
	Create(createReq string) (string, error)
}

type delivery struct {
	nc         *nats.Conn
	subjPrefix string
	queue      string
}

func NewDelivery(nc *nats.Conn, subjPrefix, queue string) Delivery {
	return &delivery{
		nc:         nc,
		subjPrefix: subjPrefix,
		queue:      queue,
	}
}

func (d *delivery) Create(createReq string) (string, error) {
	return "", nil
}
