package consumer

import (
	"github.com/shaswata56/grabbitmq/helpers"
	"github.com/shaswata56/grabbitmq/models"
	"github.com/streadway/amqp"
)

type Operation interface {
	Connect() error
	CreateChannel(channelName string) error
	Consume() (<-chan amqp.Delivery, error)
}

type Consumer struct {
	models.Amqp
}

func GetConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) CreateChannel(channelName string) (err error) {
	c.Chan, err = c.Conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")

	c.Queue, err = c.Chan.QueueDeclare(
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

func (c *Consumer) Consume() (channel <-chan amqp.Delivery, err error) {
	channel, err = c.Chan.Consume(
		c.Queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return channel, err
}

func (c *Consumer) Connect() {
	connString := helpers.GetUriMQ()
	var err error
	c.Conn, err = amqp.Dial(connString)
	helpers.FailOnError(err, "Failed to connect to RabbitMQ instance")
}

func (c *Consumer) Close() {
	c.Chan.Close()
	c.Conn.Close()
}
