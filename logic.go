package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var HTTPClient = &http.Client{}

func MessageCreatedHandler(message Message, flowiseApi string, flowiseKey string) error {
	if message.Content == "" {
		return nil
	}
	if message.ContentType != "text" {
		return nil
	}
	if message.MessageType == "outgoing" {
		return nil
	}
	if message.Conversation.AssigneeID != 0 {
		return nil
	}
	if message.Conversation.Status == "open" {
		return nil
	}

	intent, err := PredictIntentionFlowise(message.Content)
	if err != nil {
		return err
	}

	agentBots, err := ListAgentBots(strconv.Itoa(message.Account.ID))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(agentBots) == 0 {
		fmt.Println("No agent bots found")
		return errors.New("No agent bots found")
	}

	if message.Content == "/assign" || intent == "int_compra" || intent == "int_soporte" || intent == "int_devolucion" {

		// assign to agent
		resp, err := OpenConversation(strconv.Itoa(message.Account.ID), strconv.Itoa(message.Conversation.ID), agentBots[0].AccessToken)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error opening conversation", resp)
			return errors.New("Error assigning conversation to agent")
		}

		var respMessage Message
		respMessage.MessageType = "outgoing"
		respMessage.ContentType = "text"
		respMessage.Private = false
		respMessage.Account.ID = message.Account.ID
		respMessage.Conversation.ID = message.Conversation.ID
		respMessage.Content = "Tu conversación ha sido asignada a un agente.\n\nEn la brevedad se contactarán contigo para ayudarte.\n\nMuchas Gracias 😊"

		respMsg, err := SendTextMessage(respMessage, agentBots[0].AccessToken)
		if err != nil {
			return err
		}

		if respMsg.StatusCode != http.StatusOK {
			return errors.New("Error sending message")
		}

		return nil
	}

	iaResp, err := QueryFlowise(message.Content, flowiseApi, flowiseKey)
	if err != nil {
		return err
	}

	var respMessage Message
	respMessage.Content = iaResp
	respMessage.MessageType = "outgoing"
	respMessage.ContentType = "text"
	respMessage.Private = false
	respMessage.Account.ID = message.Account.ID
	respMessage.Conversation.ID = message.Conversation.ID

	fmt.Println("sending text message")

	resp, err := SendTextMessage(respMessage, agentBots[0].AccessToken)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error sending message")
		return errors.New("Error sending message")
	}

	return nil
}

func SendTextMessage(message Message, agentBotToken string) (http.Response, error) {
	var strAccID string = strconv.Itoa(message.Account.ID)
	var strConvID string = strconv.Itoa(message.Conversation.ID)
	var url string = os.Getenv("CHATWOOT_HOST") + "/api/v1/accounts/" + strAccID + "/conversations/" + strConvID + "/messages"

	fmt.Println(url, message.Account.ID, message.Conversation.ID)

	payloadBuf := new(bytes.Buffer)
	err := json.NewEncoder(payloadBuf).Encode(message)
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBuf.Bytes()))
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("api_access_token", agentBotToken)

	resp, err := HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	defer resp.Body.Close()

	return *resp, nil
}

func AssignConversationToAnAgent(accountID string, conversationID string, agentID int) (http.Response, error) {

	body := map[string]int{
		"assignee_id": agentID,
	}

	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(body); err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}

	req, err := http.NewRequest("POST", os.Getenv("CHATWOOT_HOST")+"/api/v1/accounts/"+accountID+"/conversations/"+conversationID+"/assignments", bytes.NewBuffer(payloadBuf.Bytes()))
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("api_access_token", os.Getenv("API_ACCESS_TOKEN"))

	resp, err := HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	defer resp.Body.Close()

	return *resp, nil
}

func OpenConversation(accountID string, conversationID string, agentBotToken string) (http.Response, error) {
	fmt.Println(accountID, conversationID)

	body := map[string]string{
		"status": "open",
	}

	payloadBuf := new(bytes.Buffer)
	if err := json.NewEncoder(payloadBuf).Encode(body); err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}

	req, err := http.NewRequest("POST", os.Getenv("CHATWOOT_HOST")+"/api/v1/accounts/"+accountID+"/conversations/"+conversationID+"/toggle_status", bytes.NewBuffer(payloadBuf.Bytes()))
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusBadRequest}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("api_access_token", agentBotToken)

	fmt.Println("sending request to open conversation")

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	defer resp.Body.Close()

	return *resp, nil
}

func ReturnHelloWorld() string {
	return "Hello World"
}

func ListAgentBots(accountID string) ([]AgentBot, error) {
	req, err := http.NewRequest("GET", os.Getenv("CHATWOOT_HOST")+"/api/v1/accounts/"+accountID+"/agent_bots", nil)
	if err != nil {
		return []AgentBot{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("api_access_token", os.Getenv("API_ACCESS_TOKEN"))

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return []AgentBot{}, err
	}

	defer resp.Body.Close()

	fmt.Println(resp)

	var agentBots []AgentBot
	if err := json.NewDecoder(resp.Body).Decode(&agentBots); err != nil {
		return []AgentBot{}, err
	}

	return agentBots, nil

}
