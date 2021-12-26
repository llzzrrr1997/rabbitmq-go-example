package main

import (
	"fmt"
	"github.com/llzzrrr1997/rabbitmq-go-example/work_queues"
	"time"
)

func main() {
	mq := work_queues.NewRabbitMQInstance("word_queues_durable")
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("word_queues msg:%d", i)
		mq.Publish(msg)
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}
