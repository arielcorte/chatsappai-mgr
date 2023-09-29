package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	var webhookEvent WebhookEvent

	fmt.Println(webhookEvent)

	if err := c.BindJSON(&webhookEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(webhookEvent)

	switch webhookEvent.(type) {

	case Message:
		var message Message
		if err := c.BindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(message)

		//if err := MessageCreatedHandler(message); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//	return
		//}

		var sendMessage Message = Message{
			Content:     "Hello",
			MessageType: "outgoing",
			ContentType: "text",
			Private:     false,
			Account: Account{
				ID: 1,
			},
			Conversation: Conversation{
				ID: 7,
			},
		}

		resp, err := SendTextMessage(sendMessage)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error sending message"})
			return
		}

	case Conversation:
		var conversation Conversation
		if err := c.BindJSON(&conversation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}

func SendTestHandler(c *gin.Context) {
	accountID, err := strconv.Atoi(c.Query("account_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conversationID, err := strconv.Atoi(c.Query("conversation_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sendMessage Message = Message{
		Content:     "Hello, test message sent",
		MessageType: "outgoing",
		ContentType: "text",
		Private:     false,
		Account: Account{
			ID: accountID,
		},
		Conversation: Conversation{
			ID: conversationID,
		},
	}

	resp, err := SendTextMessage(sendMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Status})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
