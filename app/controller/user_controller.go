package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rocket-api/app/constant"
	"rocket-api/app/util"
)

// 注册接受参数
type RegisterParams struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户注册
func Register(context *gin.Context) {
	var registerInfo RegisterParams
	if err := context.ShouldBind(&registerInfo); err != nil {
		respInfo := util.RespReturn(constant.RECEIVE_PARAMS_ERR, "", "参数接收错误")
		context.JSON(http.StatusOK, respInfo)
	}

	log.Println(registerInfo)
}

// 用户登录
func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "login"})
}
