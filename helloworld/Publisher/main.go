package main

import "github.com/llzzrrr1997/rabbitmq-go-example/helloworld"

func main() {
	instance := helloworld.NewRabbitMQInstance("hello_world")
	instance.Publish("hello,llzzrrr")
	instance.Close()
}
