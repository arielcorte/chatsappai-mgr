package main

import "github.com/gin-gonic/gin"

var Engine *gin.Engine

func CreateURLMappings() *gin.Engine {
	Engine = gin.Default()

	Engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	Engine.POST("/webhooks", WebhookHandler)

	Engine.POST("/send-test", SendTestHandler)
	Engine.POST("/open-conversation", AssingConversationHandler)

	return Engine
}
