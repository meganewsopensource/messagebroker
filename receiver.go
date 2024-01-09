package messagebroker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type Receiver interface {
	Receive() (<-chan amqp.Delivery, error)
}

type receiver struct {
	channel   BrokerChannel
	queueName string
}

func NewReceiver(channel BrokerChannel, queueName string) Receiver {
	return &receiver{
		channel:   channel,
		queueName: queueName,
	}
}

func (r receiver) Receive() (<-chan amqp.Delivery, error) {
	messages, err := r.channel.Consume(
		r.queueName, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
