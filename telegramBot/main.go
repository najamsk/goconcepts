package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// Update is a Telegram object that the handler receives every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object that can be found in an update.
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
}

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

// parseTelegramRequest handles incoming update from the Telegram web hook
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {
	cmdReply := ""
	// Parse incoming request
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// Sanitize input
	// var sanitizedSeed = sanitize(update.Message.Text)
	switch update.Message.Text {
	case "/ping":
		cmdReply = "pong"
	case "/about":
		cmdReply = "this is a a sample telegram bot written in golang, hosted on gcp cloud functions."
	case "/dev":
		cmdReply = "NSA"
	default:
		cmdReply = "unknown command"
	}

	// Send the punchline back to Telegram
	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, cmdReply)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	}
}

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id
func sendTextToTelegramChat(chatId int, text string) (string, error) {

	// log.Printf("Sending %s to chat_id: %d", text, chatId)
	// var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("tg") + "/sendMessage"
	response, err := http.PostForm(
		"https://api.telegram.org/bot"+os.Getenv("TG")+"/sendMessage",
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		// log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		// log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	// bodyString := string(bodyBytes)
	// log.Printf("Body of Telegram Response: %s", bodyString)

	return string(bodyBytes), nil
}
