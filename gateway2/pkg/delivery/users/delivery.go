package users

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

const (
	timeout = time.Millisecond * 500
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
	createSubj := d.subjPrefix + ".create"
	data, _ := ioutil.ReadAll(c.Request.Body)
	msg, err := d.nc.Request(createSubj, data, timeout)
	if err != nil {
		log.Printf("msg: %v, err: %v", msg, err)
	}
	log.Printf("Respuesta Subject %s", msg.Subject)
	log.Printf("Respuesta Reply %s", msg.Reply)
	log.Printf("Respuesta Data %s", msg.Data)
	c.IndentedJSON(http.StatusOK, fmt.Sprintf("%s", msg.Data))
}
