package amqp

import (
	"mocks_rabbitmq/mocks"
)

type Publisher struct {
	conn    mocks.Conn
	channel mocks.Publisher
}

func NewPublisher(conn mocks.Conn, channel mocks.Channel) (*Publisher, error) {
	var err error

	pb := Publisher{
		conn: conn,
	}

	if channel == nil {
		channel, err = conn.Channel()

		if err != nil {
			return nil, err
		}
	}

	pb.channel = channel

	return &pb, nil
}

func (pb *Publisher) Publish(exc string, route string, message []byte, opt mocks.Option) error {
	err := pb.channel.Publish(
		exc,   // publish to an exchange
		route, // routing to 0 or more queues
		message,
		opt,
	)

	return err
}
