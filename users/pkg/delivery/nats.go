package delivery

import (
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/user"
	"github.com/nats-io/nats.go"
)

type delivery struct {
	usecase user.Usecase
	ec      *nats.EncodedConn
}

func newDelivery(uc user.Usecase, ec *nats.EncodedConn) *delivery {
	return &delivery{
		usecase: uc,
		ec:      ec,
	}
}

func (d *delivery) Create(subj, reply string, user *user.User) {
	userCreated, _ := d.usecase.Create(user)

	d.ec.Publish(reply, userCreated)
}

func Subscribe(usecase user.Usecase, ec *nats.EncodedConn, subjPrefix string, queue string) error {

	delivery := newDelivery(usecase, ec)

	createSubj := subjPrefix + ".create"
	ec.QueueSubscribe(createSubj, queue, delivery.Create)

	return nil
}
