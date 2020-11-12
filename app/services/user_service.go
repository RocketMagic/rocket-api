package services

import (
	"rocket-api/app/mq"
	"rocket-api/app/util"
)

// 用户注册
func UserRegister(data map[string]interface{}) {
	registerInfoToMQ(data)
}

// 用户登录
func UserLogin() {

}

// 使用mq 消费注册
func registerInfoToMQ(data map[string]interface{}) {
	mq.RegisterPublish(string(util.JsonEncode(data)))
}
