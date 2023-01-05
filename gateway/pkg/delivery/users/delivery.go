package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

const (
	timeout = time.Second * 2
)

type Delivery interface {
	Create(c *gin.Context)
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

func (d *delivery) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create User",
	})
}
