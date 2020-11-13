package controller

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"net/http"
	"rocket-api/app/constant"
	"rocket-api/app/request"
	"rocket-api/app/services"
	"rocket-api/app/util"
	"time"
)

// 用户注册
func Register(context *gin.Context) {
	var registerInfo request.RegisterParams
	if err := context.ShouldBind(&registerInfo); err != nil {
		respInfo := util.RespReturn(constant.RECEIVE_PARAMS_ERR, "", "参数接收错误")
		context.JSON(http.StatusOK, respInfo)
		return
	}

	// 调用注册服务
	regData := map[string]interface{}{
		"Name":          registerInfo.Name,
		"Email":         registerInfo.Email,
		"Password":      util.Hash(registerInfo.Password, md5.New()),
		"LastLoginTime": time.Now().Unix(),
		"RegisterIp":    context.Request.RemoteAddr,
		"CreateTime":    time.Now().Unix(),
	}
	services.UserRegister(regData)

	respInfo := util.RespReturn(0, regData, "注册成功")
	context.JSON(http.StatusOK, respInfo)
	return
}

// 用户登录
func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "login"})
}
