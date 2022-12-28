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

func (d *delivery) Drain() {
	d.ec.Drain()
}

func Subscribe(usecase user.Usecase, ec *nats.EncodedConn, subjPrefix string, queue string) error {

	delivery := newDelivery(usecase, ec)

	createSubj := subjPrefix + ".create"
	ec.QueueSubscribe(createSubj, queue, delivery.Create)

	return nil
}

func NewDelivery(uc user.Usecase, natsURLs, subjPrefix, queue string) (*delivery, error) {
	nc, err := nats.Connect(natsURLs)
	if err != nil {
		return nil, err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	defer ec.Close()

	err = Subscribe(uc, ec, subjPrefix, queue)
	if err != nil {
		return nil, err
	}

	delivery := newDelivery(uc, ec)
	return delivery, nil
}
