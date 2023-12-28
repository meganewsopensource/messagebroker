package messagebroker

import (
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type BrokerMock struct {
	publishedMessages []string
	ReceiverError     bool
	SenderError       bool
}

func (m *BrokerMock) Receive() (<-chan amqp.Delivery, error) {
	if m.ReceiverError {
		return nil, errors.New("an error occurred when trying to receive messagebroker")
	}

	deliveryChan := make(chan amqp.Delivery, 2)

	for _, msg := range m.publishedMessages {
		msgChannel := amqp.Delivery{
			Body: []byte(fmt.Sprintf("%s", msg)),
		}
		deliveryChan <- msgChannel
	}

	close(deliveryChan)

	return deliveryChan, nil
}

func (m *BrokerMock) Send(message string) error {
	if m.SenderError {
		return errors.New("an error occurred when trying to send a messagebroker")
	}

	m.publishedMessages = append(m.publishedMessages, message)
	return nil
}
