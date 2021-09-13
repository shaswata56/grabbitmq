package models

import (
	"github.com/streadway/amqp"
)

type Amqp struct {
	Chan  *amqp.Channel
	Conn  *amqp.Connection
	Queue amqp.Queue
}
