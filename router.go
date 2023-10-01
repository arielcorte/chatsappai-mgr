package main

import "github.com/gin-gonic/gin"

var Engine *gin.Engine

func CreateURLMappings() {
	Engine = gin.Default()

	Engine.POST("/webhooks", WebhookHandler)

	Engine.POST("/send-test", SendTestHandler)
	Engine.POST("/assign-conversation", AssingConversationHandler)
}
