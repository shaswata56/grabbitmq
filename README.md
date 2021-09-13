# grabbitmq
### A high-level RabbitMQ driver for Golang.

#### Import in your project:
```sh
go get github.com/shaswata56/grabbitmq
````

#### Usage Demo:

```go
package main

import (
	"github.com/shaswata56/grabbitmq"
	"github.com/shaswata56/grabbitmq/consumer"
	"github.com/shaswata56/grabbitmq/helpers"
	"github.com/shaswata56/grabbitmq/publisher"
	"log"
)

func main() {
	grabbitmq.SetRabbitMQUri("amqps://username:password@bonobo.rmq.cloudamqp.com/random")

	pub := publisher.GetPublisher()
	pub.Connect()
	err := pub.CreateChannel("demoChannel")
	helpers.FailOnError(err, "Failed to create channel on publisher")
	err = pub.Publish([]byte("Hello RabbitMQ from demoChannel!!!"))
	helpers.FailOnError(err, "Fail to publish message")
	pub.Close()

	con := consumer.GetConsumer()
	con.Connect()
	err = con.CreateChannel("demoChannel")
	helpers.FailOnError(err, "Failed to create channel on publisher")
	messages, err := con.Consume()

	for message := range messages {
		log.Println(string(message.Body))
		if message.Body != nil {
			con.Close()
			break
		}
	}
}
```

#### Output:

```sh
2021/09/13 23:54:41 Hello RabbitMQ from demoChannel!!!
```
