package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"scrapper/scrapper"
	"github.com/joho/godotenv"
)

func getEnvVariableFromFile(key, envFile string) string {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}

	return os.Getenv(key)
}

type slackMessageAttachmentsFormat struct {
	Color string `json:"color"`
	Pretext string `json:"prefix"`
	Text string `json:"text"`
}

type slackMessageBody struct {
	Channel  string `json:"channel"`
	Text 	 string `json:"text"`
	Attachments []slackMessageAttachmentsFormat `json:"attachments"`
}

// Configure a init settings as env variable and other for launch worker
func initConfig() error {
	return nil
}

func sendSlackMessage(webhookUrl, channel, message string, attachments []slackMessageAttachmentsFormat) error {
	messageBody := slackMessageBody{
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

//Building Message for send to slack
func buildMessage(covidInfo scrapper.Item) (msg slackMessageAttachmentsFormat) {
	msg = slackMessageAttachmentsFormat{
		Color:   "#36a64f",
		Text:    fmt.Sprint("오늘까지의 누적 확진자 수는 ", covidInfo.DecideCnt, "명 입니다. :sob:"),
	}
	return
}

func main() {
	_ = initConfig()
	_ = getEnvVariableFromFile("OPEN_API_KEY", "../.env")

	fmt.Println("Start")
	todayCovidInfo := scrapper.Scrape("dump.xml")
	builtMessage := buildMessage(todayCovidInfo)
	slackWebhookUrl := getEnvVariableFromFile("SLACK_WEBHOOK_URL","../.env")
	_ = sendSlackMessage(slackWebhookUrl,"bot-test", "오늘의 코로나 알림 :mask:", []slackMessageAttachmentsFormat{builtMessage})

	fmt.Println("End")
}
