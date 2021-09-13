package grabbitmq

import "github.com/shaswata56/grabbitmq/helpers"

func SetRabbitMQUri(connectionUri string) {
	helpers.RabbitMQUri = connectionUri
}
