package main

import (
	"github.com/llzzrrr1997/rabbitmq-go-example/publisher_confirms"
)

func main() {
	mq := publisher_confirms.NewRabbitMQInstance("publisher_confirms")
	mq.Consume()
}
