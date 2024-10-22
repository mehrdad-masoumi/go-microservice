package rabbitmq

import (
	"fmt"
	"mlm/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbitmq struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func Connect(config config.Rabbitmq) (*Rabbitmq, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	))
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Rabbitmq{
		Conn:    conn,
		Channel: ch,
	}, err
}

func (r *Rabbitmq) DeclareQueue(qName string) (amqp.Queue, error) {
	q, err := r.Channel.QueueDeclare(
		qName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return q, err
	}
	return q, err
}

func (r *Rabbitmq) Publish(qName string, message string, retryCount int) error {
	headers := amqp.Table{"retry_count": retryCount}
	err := r.Channel.Publish(
		"",
		qName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Headers:     headers,
		},
	)
	return err
}

func (r *Rabbitmq) Consume(qName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(
		qName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	return msgs, err
}
