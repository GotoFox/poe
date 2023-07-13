package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/poe-openai-proxy/conf"
	"github.com/juzeon/poe-openai-proxy/poe"
	"github.com/juzeon/poe-openai-proxy/router"
	"strconv"
)

func handler(c *gin.Context) {
	conf.Setup()
	poe.Setup()
	router.Setup(c)
}

func main() {
	engine := gin.New()
	engine.GET("/*path", handler)
	engine.POST("/*path", handler)

	port := strconv.Itoa(conf.Conf.Port)
	if err := engine.Run(":" + port); err != nil {
		panic(err)
	}
}
