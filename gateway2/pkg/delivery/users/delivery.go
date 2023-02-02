package users

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/user"
	"github.com/nats-io/nats.go"
)

const (
	timeout = time.Millisecond * 500
)

type Delivery interface {
	Create(c *gin.Context)
}

type delivery struct {
	ec         *nats.EncodedConn
	subjPrefix string
	queue      string
}

func NewDelivery(ec *nats.EncodedConn, subjPrefix, queue string) Delivery {
	return &delivery{
		ec:         ec,
		subjPrefix: subjPrefix,
		queue:      queue,
	}
}

func (d *delivery) Create(c *gin.Context) {
	userCreated := &user.User{}
	createSubj := d.subjPrefix + ".create"
	data, _ := ioutil.ReadAll(c.Request.Body)
	err := d.ec.Request(createSubj, data, userCreated, timeout)
	if err != nil {
		log.Printf("err: %v", err)
	}
	c.IndentedJSON(http.StatusOK, userCreated)
}
