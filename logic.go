package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var HTTPClient = &http.Client{}

func MessageCreatedHandler(message MessageCreatedEvent, flowiseApi string, flowiseKey string) error {
	if message.Content == "" {
		return nil
	}
	if message.ContentType != "text" {
		return nil
	}
	if message.MessageType != "incoming" {
		return nil
	}
	if message.Conversation.Status == "open" {
		for _, label := range message.Conversation.Labels {
			if label == "bot" {
				return nil
			}
			if label == "manual" {
				return nil
			}
		}
	}
	for _, label := range message.Conversation.Labels {
		if label == "bot" {
			return nil
		}
	}

	intent, err := PredictIntentionFlowise(message.Content)
	if err != nil {
		return err
	}
	fmt.Println("intent", intent)

	agentBots, err := ListAgentBots(strconv.Itoa(message.Account.ID))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(agentBots) == 0 {
		fmt.Println("No agent bots found")
		return errors.New("No agent bots found")
	}

	if message.Content == "/hour" {

		loc, err := time.LoadLocation("America/Santiago")
		if err != nil {
			return err
		}

		now := time.Now().In(loc)

		resp, err := SendTextMessage(MessageCreatedEvent{
			MessageType:  "outgoing",
			ContentType:  "text",
			Private:      false,
			Account:      message.Account,
			Conversation: message.Conversation,
			Content:      now.Format("15:04") + strings.ToLower(now.Weekday().String())[0:2],
		}, agentBots[0].AccessToken)

		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return errors.New("Error sending message")
		}

		return nil
	}

	if message.Content == "/isworkhour" {
		busyCannedResponse, err := GetCannedResponseByShortCode("busy", strconv.Itoa(message.Account.ID))
		if err != nil {
			return err
		}

		_, busyWorkHours, found := strings.Cut(busyCannedResponse.ShortCode, " ")
		if !found {
			return errors.New("Error parsing work hours")
		}

		workHours, err := ParseWorkHours(busyWorkHours)
		if err != nil {
			return err
		}

		isWorkHour, diff := IsWorkHour(workHours)

		resp, err := SendTextMessage(MessageCreatedEvent{
			MessageType:  "outgoing",
			ContentType:  "text",
			Private:      false,
			Account:      message.Account,
			Conversation: message.Conversation,
			Content:      strconv.FormatBool(isWorkHour) + " " + diff.String(),
		}, agentBots[0].AccessToken)

		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return errors.New("Error sending message")
		}

		return nil
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

		busyCannedResponse, err := GetCannedResponseByShortCode("busy", strconv.Itoa(message.Account.ID))
		if err != nil {
			var respMessage MessageCreatedEvent
			respMessage.MessageType = "outgoing"
			respMessage.ContentType = "text"
			respMessage.Private = false
			respMessage.Account.ID = message.Account.ID
			respMessage.Conversation.ID = message.Conversation.ID
			respMessage.Content = "Tu conversación ha sido asignada a un agente.\n\nEn la brevedad se contactarán contigo para ayudarte.\n\nMuchas Gracias 😊"

			respMsg, new_err := SendTextMessage(respMessage, agentBots[0].AccessToken)
			if new_err != nil {
				return err
			}

			if respMsg.StatusCode != http.StatusOK {
				return errors.New("Error sending message")
			}
		}

		if busyCannedResponse.Content != "" {

			cannedResponseWorkHours, err := ParseWorkHours(busyCannedResponse.ShortCode)
			if err != nil {
				return err
			}

			isWorkHour, diff := IsWorkHour(cannedResponseWorkHours)

			if isWorkHour {
				var respMessage MessageCreatedEvent
				respMessage.MessageType = "outgoing"
				respMessage.ContentType = "text"
				respMessage.Private = false
				respMessage.Account.ID = message.Account.ID
				respMessage.Conversation.ID = message.Conversation.ID
				if diff <= time.Duration(45*time.Minute) {
					busyWarn, err := GetCannedResponseByShortCode("busy-warn", strconv.Itoa(message.Account.ID))
					if err != nil {
						return err
					}
					if busyWarn.Content == "" {
						respMessage.Content = busyWarn.Content
					}
				} else {
					busyWait, err := GetCannedResponseByShortCode("busy-wait", strconv.Itoa(message.Account.ID))
					if err != nil {
						return err
					}
					if busyWait.Content == "" {
						respMessage.Content = busyWait.Content
					}
				}

				if respMessage.Content == "" {
					return errors.New("Error sending message")
				}
				respMsg, new_err := SendTextMessage(respMessage, agentBots[0].AccessToken)
				if new_err != nil {
					return err
				}

				if respMsg.StatusCode != http.StatusOK {
					return errors.New("Error sending message")
				}

				return nil
			}

			var respMessage MessageCreatedEvent
			respMessage.MessageType = "outgoing"
			respMessage.ContentType = "text"
			respMessage.Private = false
			respMessage.Account.ID = message.Account.ID
			respMessage.Conversation.ID = message.Conversation.ID
			respMessage.Content = busyCannedResponse.Content

			respMsg, new_err := SendTextMessage(respMessage, agentBots[0].AccessToken)
			if new_err != nil {
				return err
			}

			if respMsg.StatusCode != http.StatusOK {
				return errors.New("Error sending message")
			}
		}

		return nil
	}

	iaResp, err := QueryFlowise(message.Content, strconv.Itoa(message.Sender.ID), message.Inbox.Name, flowiseApi, flowiseKey)
	if err != nil {
		return err
	}

	var respMessage MessageCreatedEvent
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

func SendTextMessage(message MessageCreatedEvent, agentBotToken string) (http.Response, error) {
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

	var agentBots []AgentBot
	if err := json.NewDecoder(resp.Body).Decode(&agentBots); err != nil {
		return []AgentBot{}, err
	}

	return agentBots, nil

}

func ConversationStatusChangedHandler(conversation ConversationStatusChangedEvent, flowiseUrl string, flowiseKey string) error {
	return nil
}

func GetAllCannedResponses(accountID string) ([]CannedResponse, error) {
	req, err := http.NewRequest("GET", os.Getenv("CHATWOOT_HOST")+"/api/v1/accounts/"+accountID+"/canned_responses", nil)
	if err != nil {
		return []CannedResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("api_access_token", os.Getenv("API_ACCESS_TOKEN"))

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return []CannedResponse{}, err
	}

	defer resp.Body.Close()

	var cannedResponses []CannedResponse
	if err := json.NewDecoder(resp.Body).Decode(&cannedResponses); err != nil {
		return []CannedResponse{}, err
	}

	return cannedResponses, nil
}

func GetCannedResponseByShortCode(shortCode string, accountID string) (CannedResponse, error) {
	cannedResponses, err := GetAllCannedResponses(accountID)
	if err != nil {
		return CannedResponse{}, err
	}

	for _, cannedResponse := range cannedResponses {
		if strings.HasPrefix(cannedResponse.ShortCode, shortCode) {
			return cannedResponse, nil
		}
	}

	return CannedResponse{}, errors.New("Canned response not found")
}

func ParseWorkHours(busyWorkHours string) ([]WorkHours, error) {

	weekMap := map[string]int{
		"mon": 0,
		"tue": 1,
		"wed": 2,
		"thu": 3,
		"fri": 4,
		"sat": 5,
		"sun": 6,
	}

	weekArray := [7]string{
		"mon",
		"tue",
		"wed",
		"thu",
		"fri",
		"sat",
		"sun",
	}

	englishToSpanish := map[string]string{
		"mon": "lun",
		"tue": "mar",
		"wed": "mie",
		"thu": "jue",
		"fri": "vie",
		"sat": "sab",
		"sun": "dom",
	}

	//busyWorkHours format: "busy Mo-Fr:09:00-17:00;Sa-Su:10:00-14:00"

	//translate spanish to english
	for key, value := range englishToSpanish {
		busyWorkHours = strings.ReplaceAll(busyWorkHours, value, key)
	}

	fmt.Println(busyWorkHours)

	loc, err := time.LoadLocation("America/Santiago")
	if err != nil {
		return []WorkHours{}, err
	}

	dayHourPairs := strings.Split(busyWorkHours, ";")

	var workHours []WorkHours

	for _, pair := range dayHourPairs {
		day, hour, found := strings.Cut(pair, ":")

		if !found {
			return []WorkHours{}, errors.New("Error parsing work hours")
		}

		startHour, endHour, foundHours := strings.Cut(hour, "-")
		if !foundHours {
			return []WorkHours{}, errors.New("Error parsing work hours")
		}

		start, err := time.ParseInLocation("15:04", startHour, loc)
		if err != nil {
			return []WorkHours{}, err
		}

		end, err := time.ParseInLocation("15:04", endHour, loc)
		if err != nil {
			return []WorkHours{}, err
		}

		startDay, endDay, foundDays := strings.Cut(day, "-")

		if foundDays {
			if weekMap[startDay] > weekMap[endDay] {
				return []WorkHours{}, errors.New("Error parsing work days")
			}

			for i := weekMap[startDay]; i <= weekMap[endDay]; i++ {
				workHours = append(workHours, WorkHours{
					Day:   weekArray[i],
					Start: start,
					End:   end,
				})
			}
		}

		if !foundDays {
			workHours = append(workHours, WorkHours{
				Day:   day,
				Start: start,
				End:   end,
			})
		}
	}

	//remove duplicate days

	var uniqueWorkHours []WorkHours

	for _, workHour := range workHours {
		found := false
		for _, uniqueWorkHour := range uniqueWorkHours {
			if uniqueWorkHour.Day == workHour.Day {
				found = true
			}
		}
		if !found {
			uniqueWorkHours = append(uniqueWorkHours, workHour)
		}
	}

	return uniqueWorkHours, nil

}

func IsWorkHour(workHours []WorkHours) (bool, time.Duration) {
	loc, err := time.LoadLocation("America/Santiago")
	if err != nil {
		return false, 0
	}

	now := time.Now().In(loc)

	for _, workHour := range workHours {
		fmt.Println("day", workHour.Day)
		if workHour.Day == strings.ToLower(now.Weekday().String())[0:len(workHour.Day)] {
			newStart := time.Date(now.Year(), now.Month(), now.Day(), workHour.Start.Hour(), workHour.Start.Minute(), 0, 0, loc)
			newEnd := time.Date(now.Year(), now.Month(), now.Day(), workHour.End.Hour(), workHour.End.Minute(), 0, 0, loc)
			fmt.Println(newStart, "|", now, "|", newEnd)
			fmt.Println(now.After(newStart), now.Before(newEnd))
			if now.After(newStart) && now.Before(newEnd) {
				return true, newEnd.Sub(now).Truncate(time.Minute)
			}
		}
	}

	return false, 0
}
