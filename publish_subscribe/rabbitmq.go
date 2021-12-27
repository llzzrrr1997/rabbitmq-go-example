package publish_subscribe

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// MQURL 连接信息amqp://lzr:123qwe@127.0.0.1:5672/hello_world 这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://lzr:123qwe@127.0.0.1:5672/example"

// RabbitMQ rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	MqURL string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqURL: MQURL}
}

// Close 关闭mq
func (r *RabbitMQ) Close() {
	_ = r.channel.Close() //先关闭信道
	_ = r.conn.Close()    //再关闭链接
}

//错误处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
	}
}

func NewRabbitMQInstance(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqURL)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

// Publish 生产消息
func (r *RabbitMQ) Publish(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("ExchangeDeclare err:%s", err)
		return
	}
	//调用channel 发送消息到队列中
	_ = r.channel.Publish(
		r.Exchange,
		"",
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// Consume 消费消息
func (r *RabbitMQ) Consume() {
	//尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("ExchangeDeclare err:%s", err)
		return
	}
	//申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		"", //随机产生队列名
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Printf("QueueDeclare err:%s", err)
		return
	}
	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("QueueBind err:%s", err)
		return
	}
	//接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Connection中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//队列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Printf("Consume err:%s", err)
	}
	forever := make(chan struct{})
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理
			msg := d.Body
			log.Printf("start received a message: %s", msg)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
