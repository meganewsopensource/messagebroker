package messagebroker

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MockChannel struct {
	messages              []string
	errorConsumingMessage bool
}

func (mc *MockChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
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
