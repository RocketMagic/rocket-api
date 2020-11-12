package mq

import (
	"fmt"
	"log"
	"rocket-api/app/entity"
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
	mq.ReceiveSub(func(body []byte) {
		var user entity.Users
		util.JsonDecode(body, &user)
		fmt.Println(user)
		result := InsertUser(user)
		if result {
			log.Printf("【注册用户写入数据失败】\n")
		}
	})
}

// 新增用户到数据库
func InsertUser(user entity.Users) bool {
	return util.DB.NewRecord(user)
}
