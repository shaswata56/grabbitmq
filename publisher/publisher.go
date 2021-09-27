package publisher

import (
	"github.com/shaswata56/grabbitmq/helpers"
	"github.com/shaswata56/grabbitmq/models"
	"github.com/streadway/amqp"
)

type Operation interface {
	Connect() error
	Publish(blob []byte) error
	Close()
}

type Publisher struct {
	models.Amqp
}

func GetPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Publish(exchange string, blob []byte) (err error) {
	err = p.Chan.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/vnd.masstransit+json",
			Body:        blob,
		})
	return err
}

func (p *Publisher) Connect() {
	var err error
	connString := helpers.GetUriMQ()
	p.Conn, err = amqp.Dial(connString)
	helpers.FailOnError(err, "Failed to connect to RabbitMQ instance")
	p.Chan, err = p.Conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
}

func (p *Publisher) Close() {
	_ = p.Chan.Close()
	_ = p.Conn.Close()
}
