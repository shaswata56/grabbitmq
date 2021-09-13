package helpers

import (
	"log"
)

var RabbitMQUri string

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetUriMQ() string {
	return RabbitMQUri
}
