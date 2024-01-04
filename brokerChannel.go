package messagebroker

import amqp "github.com/rabbitmq/amqp091-go"

type BrokerChannel interface {
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
}

type RealChannel struct {
	*amqp.Channel
}

func (rc *RealChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return rc.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}
