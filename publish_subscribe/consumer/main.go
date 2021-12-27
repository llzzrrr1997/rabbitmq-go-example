package main

import "github.com/llzzrrr1997/rabbitmq-go-example/publish_subscribe"

func main() {
	mq := publish_subscribe.NewRabbitMQInstance("publish_subscribe")
	mq.Consume()
}
