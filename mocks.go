package rabbitmq_core

import (
	"bou.ke/monkey"
	"github.com/streadway/amqp"
	"mocks_rabbitmq/mocks/amqptest"
	"reflect"
)

func MockRabbitMQFunc(mockConn *amqptest.Conn) {
	// mock channel
	var c *amqp.Channel
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Publish", func(channel *amqp.Channel, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
		cha, err := mockConn.Channel()
		if err != nil {
			return err
		}
		err = cha.Publish(exchange, key, msg.Body, nil)
		if err != nil {
			return err
		}
		return nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "ExchangeDeclare", func (channel *amqp.Channel, name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
		a, err := mockConn.Channel()
		if err != nil {
			return err
		}
		err = a.ExchangeDeclare(name, kind, nil)
		if err != nil {
			return err
		}
		return nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "QueueDeclare", func (channel *amqp.Channel, name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
		a, err := mockConn.Channel()
		if err != nil {
			return amqp.Queue{}, err
		}
		queue, err := a.QueueDeclare(name, nil)
		return amqp.Queue{
			Name: queue.Name(),
			Messages: queue.Messages(),
			Consumers: queue.Consumers(),
		}, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "QueueBind", func (channel *amqp.Channel, name, key, exchange string, noWait bool, args amqp.Table) error {
		a, err := mockConn.Channel()
		if err != nil {
			return err
		}
		err = a.QueueBind(name, key, exchange, nil)
		if err != nil {
			return err
		}
		return nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Consume", func(channel *amqp.Channel, queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
		a, err := mockConn.Channel()
		if err != nil {
			return nil, err
		}
		cha, err := a.Consume(queue, consumer, nil)
		if err != nil {
			return nil, err
		}
		tempCh := make(chan amqp.Delivery, 0)
		go func() {
			for {
				select {
				case info := <- cha:
					if info != nil {
						tempinfo := amqp.Delivery{
							Body: info.Body(),
						}
						tempCh <- tempinfo
					}
				}
			}
		}()
		return tempCh, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Close", func(channel *amqp.Channel) error  {
		return nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Qos", func (channel *amqp.Channel, prefetchCount, prefetchSize int, global bool) error {
		a, err := mockConn.Channel()
		if err != nil {
			return err
		}

		err = a.Qos(prefetchCount, prefetchSize, global)
		if err != nil {
			return err
		}
		return nil
	})

	// mock amqp
	monkey.Patch(amqp.Dial, func(url string) (*amqp.Connection, error)  {
		return nil, nil
	})

	//mock connection
	var ca *amqp.Connection
	monkey.PatchInstanceMethod(reflect.TypeOf(ca), "Channel", func (connection *amqp.Connection) (*amqp.Channel, error) {
		_, err := mockConn.Channel()
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(ca), "Close", func(connection *amqp.Connection) error {
		return nil
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(ca), "IsClosed", func(connection *amqp.Connection) bool {
		return true
	})

}

func UnMockRabbitMQFunc() {
	monkey.UnpatchAll()
}