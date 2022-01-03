package main

import "github.com/llzzrrr1997/rabbitmq-go-example/routing"

func main() {
	forever := make(chan struct{})
	mq1 := routing.NewRabbitMQInstance("routing", "key_one")
	mq2 := routing.NewRabbitMQInstance("routing", "key_two")
	go func() {
		mq1.Consume()
	}()
	go func() {
		mq2.Consume()
	}()
	<-forever
}
