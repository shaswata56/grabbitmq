package consumer

import (
	"github.com/shaswata56/grabbitmq/helpers"
	"github.com/shaswata56/grabbitmq/models"
	"github.com/streadway/amqp"
)

type Operation interface {
	Connect()
	Consume(queue string) (<-chan amqp.Delivery, error)
	Close()
}

type Consumer struct {
	models.Amqp
}

func GetConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Consume(queue string) (channel <-chan amqp.Delivery, err error) {
	channel, err = c.Chan.Consume(
		queue,
		"orbitax-rabbitmq-golang",
		true,
		false,
		false,
		false,
		nil,
	)
	return channel, err
}

func (c *Consumer) Connect() {
	var err error
	connString := helpers.GetUriMQ()
	c.Conn, err = amqp.Dial(connString)
	helpers.FailOnError(err, "Failed to connect to RabbitMQ instance")
	c.Chan, err = c.Conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
}

func (c *Consumer) Close() {
	_ = c.Chan.Close()
	_ = c.Conn.Close()
}
