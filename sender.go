package messageBroker

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Sender interface {
	Send(message string) error
}

type sender struct {
	channel   *amqp.Channel
	queueName string
	ctx       context.Context
}

func NewSender(channel *amqp.Channel, queueName string, ctx context.Context) Sender {
	return &sender{
		channel:   channel,
		queueName: queueName,
		ctx:       ctx,
	}
}

func (s sender) Send(message string) error {
	body := message
	err := s.channel.PublishWithContext(s.ctx,
		"",          // exchange
		s.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Println(message)
	return err
}
