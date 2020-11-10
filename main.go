package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"rocket-api/app/mq"
	"rocket-api/app/routers/api"
	"rocket-api/app/util"
)

func main() {
	go mq.ConsumerStart() // mq消费者启动

	router := gin.Default()
	httpServer := api.SetupRouter(router)
	log.Println(httpServer.Run(fmt.Sprintf("%s:%s", util.GetConfig("server.host"), util.GetConfig("server.port"))))
}
