package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scrapper/scrapper"
	"github.com/joho/godotenv"
	"time"
)


/*

1. 일단 워커 구조 떄려ㅈ치고



 */
func getEnvVariableFromFile(key, envFile string) string {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}

	return os.Getenv(key)
}

type SlackWebhookBody struct {
	Channel  string `json:"channel"`
	Text 	 string `json:"text"`
}


// Configure a init settings as env variable and other for launch worker
func initConfig() error {
	return nil
}

func slackMessage (webhookUrl, message, channel string) error {
	payload := map[string]string{
		"channel" : channel,
		"text" 	  : message,
	}

	http.Post(webhooUrl, "application/json", )
	payload

	return nil
}

func buildSlackMessage(covidInfo Item) (message string){
	//Building Message for send to slack
	message = fmt.Sprintln("오늘의 코로나 확진자 수 :sob:")
	message = fmt.Sprintln(message, covidInfo.DecideCnt, "명")
	return
}

func main() {
	_ = initConfig()
	_ = getEnvVariableFromFile("OPEN_API_KEY", "../.env")

	fmt.Println("Start")

	todayCovidInfo := scrapper.Scrape("dump.xml")
	builtMessage := buildSlackMessage(todayCovidInfo)

	slackWebhookUrl = getEnvVariableFromFile("SLACK_WEBHOOK_URL","../.env")
	_ = slackMessage(url, builtMessage, channel)

	fmt.Println("End")
}
