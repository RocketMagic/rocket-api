package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rocket-api/app/routers/api"
)

func main() {
	router := gin.Default()
	httpServer := api.SetupRouter(router)
	log.Println(httpServer.Run(":9001"))
}
