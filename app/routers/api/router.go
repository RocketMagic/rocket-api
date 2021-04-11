package api

import (
	"github.com/gin-gonic/gin"
	"rocket-api/app/controller"
)

// 初始化路由
func SetupRouter(router *gin.Engine) *gin.Engine {

	// 用户
	router.Group("/user")
	{
		router.POST("/register", controller.Register)
		router.POST("/login", controller.Login)
	}

	return router
}
