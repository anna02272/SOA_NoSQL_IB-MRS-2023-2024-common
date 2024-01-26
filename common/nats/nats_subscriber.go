package nats

import (
	"acc-service/common/saga"
	"github.com/nats-io/nats.go"
	"log"
)

type Subscriber struct {
	conn       *nats.EncodedConn
	subject    string
	queueGroup string
}

func NewNATSSubscriber(host, port, user, password, subject, queueGroup string) (saga.Subscriber, error) {
	conn, err := getConnection(host, port, user, password)
	if err != nil {
		log.Fatal("Error connecting to NATS subscriber:", err)
		return nil, err
	}
	encConn, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		conn:       encConn,
		subject:    subject,
		queueGroup: queueGroup,
	}, nil
}

func (s *Subscriber) Subscribe(handler interface{}) error {
	_, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, handler)
	if err != nil {
		return err
	}
	return nil
}