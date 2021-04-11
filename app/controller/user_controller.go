package controller

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"net/http"
	"rocket-api/app/constant"
	"rocket-api/app/entity"
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

	// 校验邮箱格式
	if ok := util.CheckEmail(registerInfo.Email); !ok {
		respInfo := util.RespReturn(constant.RECEIVE_PARAMS_ERR, "", "邮箱格式有误")
		context.JSON(http.StatusOK, respInfo)
		return
	}

	var user entity.Users
	if res := util.DB.Where("email = ?", registerInfo.Email).First(&user); res.Error == nil {
		// 邮箱已注册
		respInfo := util.RespReturn(constant.NO_KNOW_ERR, "", "该邮箱已注册")
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
	var userInfo request.RegisterParams
	if err := context.ShouldBind(&userInfo); err != nil {
		respInfo := util.RespReturn(constant.RECEIVE_PARAMS_ERR, "", "参数接收错误")
		context.JSON(http.StatusOK, respInfo)
		return
	}

	res := services.UserLogin(userInfo)
	context.JSON(http.StatusOK, res)
	return
}
