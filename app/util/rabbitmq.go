package util

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

// 定义RabbitMQ对象
type RabbitMQ struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	queueName    string // 队列名称
	routingKey   string // key名称
	exchangeName string // 交换机名称
	exchangeType string // 交换机类型
	mu           sync.RWMutex
}

// 消费回调
type callBack func(body []byte)

// 初始化mq
func NewRabbitMQ(queueName string, routingKey string, exchangeName string, exchangeType string) *RabbitMQ {
	var err error
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/", GetConfig("rabbitmq.user"), GetConfig("rabbitmq.password"), GetConfig("rabbitmq.host"), GetConfig("rabbitmq.port"))
	mqConn, err := amqp.Dial(RabbitUrl)
	if err != nil {
		log.Printf("MQ打开链接失败:%s \n", err)
	}
	mqChan, err := mqConn.Channel()
	if err != nil {
		log.Printf("MQ打开管道失败:%s \n", err)
	}

	return &RabbitMQ{
		connection:   mqConn,
		channel:      mqChan,
		queueName:    queueName,
		routingKey:   routingKey,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
	}
}

// 关闭RabbitMQ连接
func (r *RabbitMQ) mqClose() {
	// 先关闭管道,再关闭链接
	err := r.channel.Close()
	if err != nil {
		log.Printf("MQ管道关闭失败:%s \n", err)
	}
	err = r.connection.Close()
	if err != nil {
		log.Printf("MQ链接关闭失败:%s \n", err)
	}
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.exchangeName,
		r.exchangeType,
		true,
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an excha"+
		"nge")

	//2.发送消息
	err = r.channel.Publish(
		r.exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(message),
		})
}

//订阅模式消费端代码
func (r *RabbitMQ) ReceiveSub(consumeCallBack callBack) {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.exchangeName,
		//交换机类型
		r.exchangeType,
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exch"+
		"ange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.exchangeName,
		false,
		nil)
	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range messges {
			log.Printf("Received a message: %s", d.Body)
			consumeCallBack(d.Body)
		}
	}()

	<-forever
}
