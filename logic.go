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

func MessageCreatedHandler(message Message) error {
	if message.Content == "" {
		return nil
	}
	if message.ContentType != "text" {
		return nil
	}
	if message.MessageType == "outgoing" {
		return nil
	}

	//TODO: get if conversation is manual
	var isManual bool = false

	if isManual {
		return nil
	}

	iaResp, err := QueryFlowise(message.Content)
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

	resp, err := SendTextMessage(respMessage)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error sending message")
		return errors.New("Error sending message")
	}

	return nil
}

func SendTextMessage(message Message) (http.Response, error) {
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
	req.Header.Add("api_access_token", os.Getenv("API_ACCESS_TOKEN"))

	resp, err := HTTPClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return http.Response{StatusCode: http.StatusInternalServerError}, err
	}

	defer resp.Body.Close()

	return *resp, nil
}
