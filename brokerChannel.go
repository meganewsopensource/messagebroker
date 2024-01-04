package messagebroker

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type BrokerChannel interface {
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

type RealChannel struct {
	*amqp.Channel
}

func (rc *RealChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	return rc.Channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

func (rc *RealChannel) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return rc.PublishWithContext(ctx, exchange, key, mandatory, immediate, msg)
}
