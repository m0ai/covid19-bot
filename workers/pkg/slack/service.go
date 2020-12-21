package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MessageAttachmentsFormat struct {
	Color string `json:"color"`
	Pretext string `json:"prefix"`
	Text string `json:"text"`
}

type MessageBody struct {
	Channel  string                        `json:"channel"`
	Text 	 string                        `json:"text"`
	Attachments []MessageAttachmentsFormat `json:"attachments"`
}

func SendSlackMessage(webhookUrl, channel, message string, attachments []MessageAttachmentsFormat) error {
	messageBody := MessageBody{
		Channel: channel,
		Text:    message,
		Attachments: attachments,
	}
	jsonMessage, err := json.Marshal(messageBody)
	fmt.Println(string(jsonMessage))

	if err != nil {
		log.Fatalln("Failed a json encode", messageBody)
	}
	buffer := bytes.NewBuffer(jsonMessage)
	resp, err := http.Post(webhookUrl, "application/json", buffer)
	if err != nil{
		log.Fatalln("Oh.. failure a send slack message :(", err, webhookUrl)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if !bytes.Equal(body, []byte("ok")) {
		log.Fatalln("Isn't not successful to send slack message, ", string(body))
	}

	return nil
}
