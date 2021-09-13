package publisher

import (
	"github.com/shaswata56/grabbitmq/helpers"
	"github.com/shaswata56/grabbitmq/models"
	"github.com/streadway/amqp"
)

type Operation interface {
	Connect() error
	CreateChannel(channelName string) error
	Publish(blob []byte) error
	Close()
}

type Publisher struct {
	models.Amqp
}

func GetPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) CreateChannel(channelName string) (err error) {
	p.Chan, err = p.Conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")

	p.Queue, err = p.Chan.QueueDeclare(
		channelName,
		false,
		false,
		false,
		false,
		nil,
	)
	helpers.FailOnError(err, "Failed to declare a Queue Channel")
	return err
}

func (p *Publisher) Publish(blob []byte) (err error) {
	err = p.Chan.Publish(
		"",
		p.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        blob,
		})
	return err
}

func (p *Publisher) Connect() {
	connString := helpers.GetUriMQ()
	var err error
	p.Conn, err = amqp.Dial(connString)
	helpers.FailOnError(err, "Failed to connect to RabbitMQ instance")
}

func (p *Publisher) Close() {
	p.Chan.Close()
	p.Conn.Close()
}
