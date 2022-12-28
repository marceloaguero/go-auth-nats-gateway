package users

import (
	"time"
)

const (
	timeout = time.Second * 2
)

type Delivery interface {
	Create(createReq string) (string, error)
}

type delivery struct {
	subjPrefix string
	queue      string
}

func NewDelivery(subjPrefix, queue string) Delivery {
	return &delivery{
		subjPrefix: subjPrefix,
		queue:      queue,
	}
}

func (d *delivery) Create(createReq string) (string, error) {
	return "", nil
}
