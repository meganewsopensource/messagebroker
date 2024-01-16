package messagebroker

import (
	"context"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type mockChannel struct {
	messages              []string
	errorConsumingMessage bool
	errorSendingMessage   bool
}

func (mc *mockChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	if mc.errorConsumingMessage {
		return nil, errors.New("an error occurred while trying to consume queue")
	}

	deliveries := make(chan amqp.Delivery, 2)

	for _, msg := range mc.messages {
		msgChannel := amqp.Delivery{
			Body: []byte(fmt.Sprintf("%s", msg)),
		}
		deliveries <- msgChannel
	}

	return deliveries, nil
}

func (mc *mockChannel) PublishWithContext(ctx context.Context, exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	if mc.errorSendingMessage {
		return errors.New("an error occurred while trying to send a message")
	}
	mc.messages = append(mc.messages, string(msg.Body))
	return nil
}
