package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func WebhookHandler(c *gin.Context) {
	var webhookEvent WebhookEvent

	raw, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var buf bytes.Buffer

	if err := json.Compact(&buf, raw); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := buf.Bytes()

	if err := json.Unmarshal(data, &webhookEvent); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch webhookEvent.Event {

	case "message_created":

		var message Message
		if err := json.Unmarshal(data, &message); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("message received", message.Content)

		if err := MessageCreatedHandler(message); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	case "conversation_created":
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

func AssingConversationHandler(c *gin.Context) {
	accountID := "1"
	conversationID := "16"
	assigneeID := 3

	resp, err := AssignConversationToAnAgent(accountID, conversationID, assigneeID)
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
