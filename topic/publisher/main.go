package main

import (
	"fmt"
	"github.com/llzzrrr1997/rabbitmq-go-example/topic"
	"time"
)

func main() {
	mq1 := topic.NewRabbitMQInstance("topic", "key.one")
	mq2 := topic.NewRabbitMQInstance("topic", "key.two")
	for i := 0; i < 10; i++ {
		mq1.Publish(fmt.Sprintf("topic msg_one:%d", i))
		mq2.Publish(fmt.Sprintf("topic msg_two:%d", i))
		fmt.Println("index:", i)
		time.Sleep(1 * time.Second)
	}
}
