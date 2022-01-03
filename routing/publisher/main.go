package main

import (
	"fmt"
	"github.com/llzzrrr1997/rabbitmq-go-example/routing"
	"time"
)

func main() {
	mq1 := routing.NewRabbitMQInstance("routing", "key_one")
	mq2 := routing.NewRabbitMQInstance("routing", "key_two")
	for i := 0; i < 10; i++ {
		mq1.Publish(fmt.Sprintf("routing msg_one:%d", i))
		mq2.Publish(fmt.Sprintf("routing msg_two:%d", i))
		fmt.Println("index:", i)
		time.Sleep(1 * time.Second)
	}
}
