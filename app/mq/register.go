package mq

import (
	"log"
	"rocket-api/app/util"
)

const (
	QUEUE_NAME    = "rocket_register_queue"
	ROUTING_KEY   = "rocket_register_route"
	EXCHANGE_NAME = "rocket"
	EXCHANGE_TYPE = "direct"
)

var mq = util.NewRabbitMQ(QUEUE_NAME, ROUTING_KEY, EXCHANGE_NAME, EXCHANGE_TYPE)

// 发布消息
func RegisterPublish(msg string) {
	mq.PublishPub(msg)
}

// 消费
func RegisterConsume() {
	mq.ReceiveSub(func(body string) {
		log.Println("消费消息", body)
	})
}
