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
	"covid-19-alert-to-slack/common/slackUtil"

)

// Configure a init settings as env variable and other for launch worker
func initConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	return nil
}

//Building Message for send to slack
func buildMessage(covidInfo scrapper.Item) (msg slackUtil.SlackMessageAttachmentsFormat) {
	msg = slackUtil.SlackMessageAttachmentsFormat{
		Color:   "#36a64f",
		Text:    fmt.Sprint("오늘까지의 누적 확진자 수는 ", covidInfo.DecideCnt, "명 입니다. :sob:"),
	}
	return
}

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalln("Error while initializing config")
	}

	fmt.Println("Start")

	todayCovidInfo := scrapper.Scrape(os.Getenv("OPEN_API_KEY"))
	builtMessage := buildMessage(todayCovidInfo)
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	_ = slackUtil.SendSlackMessage(slackWebhookUrl,"bot-test", "오늘의 코로나 알림 :mask:", []slackUtil.SlackMessageAttachmentsFormat{builtMessage})

	fmt.Println("End")
}
