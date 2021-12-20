package server

import "mocks_rabbitmq/mocks"

const (
	QueueMaxLen = 2 << 8
)

type Queue struct {
	name string
	data chan mocks.Delivery
}

func NewQueue(name string) *Queue {
	return &Queue{
		name: name,
		data: make(chan mocks.Delivery, QueueMaxLen),
	}
}

func (q *Queue) Consumers() int {
	return 0
}

func (q *Queue) Name() string {
	return q.name
}

func (q *Queue) Messages() int {
	return 0
}
