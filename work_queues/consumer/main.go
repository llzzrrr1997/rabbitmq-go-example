package main

import "github.com/llzzrrr1997/rabbitmq-go-example/work_queues"

func main() {
	mq := work_queues.NewRabbitMQInstance("word_queues_durable")
	mq.Consume()
}
