package messageBroker

import amqp "github.com/rabbitmq/amqp091-go"

type MessageBroker interface {
	CreateChannel() (*amqp.Channel, error)
}

type connectionMessageBroker struct {
	connection *amqp.Connection
}

func (c connectionMessageBroker) CreateChannel() (*amqp.Channel, error) {
	return c.connection.Channel()
}

func NewMessageBroker(connection string) (MessageBroker, error) {
	conn, err := amqp.Dial(connection)
	if err != nil {
		return nil, err
	}
	return connectionMessageBroker{
		connection: conn,
	}, nil
}
