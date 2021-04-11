package services

import (
	"crypto/md5"
	"log"
	"net/http"
	"rocket-api/app/constant"
	"rocket-api/app/entity"
	"rocket-api/app/mq"
	"rocket-api/app/request"
	"rocket-api/app/util"
	"strconv"
	"time"
)

// 用户注册
func UserRegister(data map[string]interface{}) {
	registerInfoToMQ(data)
}

// 用户登录
func UserLogin(user request.RegisterParams) map[string]interface{} {
	var data entity.Users
	if result := util.DB.Where("email = ?", user.Email).First(&data); result.Error != nil {
		// 找不到该用户
		return util.RespReturn(constant.RECEIVE_PARAMS_ERR, "", "用户不存在")
	}

	if data.Password != util.Hash(user.Password, md5.New()) {
		return util.RespReturn(constant.LOGIN_ERR, "", "密码错误")
	}

	// 过期时间三十分钟
	token := generateToken(data.Email)
	redisCli, err := util.RedisCli()
	if err != nil {
		log.Printf("redis error : %v\n", err)
		return util.RespReturn(constant.NO_KNOW_ERR, "", "服务器故障")
	}
	if err := redisCli.Set(token, string(util.JsonEncode(data)), time.Second*30).Err(); err != nil {
		log.Printf("redis error : %v\n", err)
		return util.RespReturn(constant.NO_KNOW_ERR, "", "服务器故障")
	}

	// success
	result := map[string]interface{}{
		"token": token,
		"info":  data,
	}

	return util.RespReturn(http.StatusOK, result, "登录成功")
}

func generateToken(email string) string {
	timestamp := time.Now().Unix()
	timeStr := strconv.FormatInt(timestamp, 10)
	start := len([]rune(timeStr)) - 5

	pre := util.StrSplice(email, -1, 5)
	next := util.StrSplice(timeStr, start, -1)

	return util.Hash(pre+next, md5.New())
}

// 使用mq 消费注册
func registerInfoToMQ(data map[string]interface{}) {
	mq.RegisterPublish(string(util.JsonEncode(data)))
}
