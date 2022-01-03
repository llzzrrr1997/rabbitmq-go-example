package main

import (
	"fmt"
	"github.com/llzzrrr1997/rabbitmq-go-example/publisher_confirms"
	"time"
)

func main() {
	mq := publisher_confirms.NewRabbitMQInstance("publisher_confirms")
	start := time.Now().UnixMilli()
	//mq.PublishSingle()
	end := time.Now().UnixMilli()
	fmt.Println("PublishSingle time:", end-start)
	start = time.Now().UnixMilli()
	//mq.PublishBatch()
	end = time.Now().UnixMilli()
	fmt.Println("PublishBatch time:", end-start)
	start = time.Now().UnixMilli()
	mq.PublishAsync()
	end = time.Now().UnixMilli()
	fmt.Println("PublishAsync time:", end-start)
	time.Sleep(3 * time.Second)
}
