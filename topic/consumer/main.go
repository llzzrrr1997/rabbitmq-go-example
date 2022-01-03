package main

import (
	"github.com/llzzrrr1997/rabbitmq-go-example/topic"
)

func main() {
	forever := make(chan struct{})
	mq1 := topic.NewRabbitMQInstance("topic", "#")
	mq2 := topic.NewRabbitMQInstance("topic", "key.two")
	go func() {
		mq1.Consume()
	}()
	go func() {
		mq2.Consume()
	}()
	<-forever
}
