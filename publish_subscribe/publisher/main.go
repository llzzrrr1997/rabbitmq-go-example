package main

import (
	"fmt"
	"github.com/llzzrrr1997/rabbitmq-go-example/publish_subscribe"
	"time"
)

func main() {
	mq := publish_subscribe.NewRabbitMQInstance("publish_subscribe")
	for i := 1; i <= 15; i++ {
		msg := fmt.Sprintf("publish_subscribe msg:%d", i)
		mq.Publish(msg)
		time.Sleep(1 * time.Second)
		fmt.Println(msg)
	}
}
