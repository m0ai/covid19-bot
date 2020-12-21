package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
	"os"
	"scrapper/internal/entity"
	scrape "scrapper/pkg/scrapper"
	slackUtil "scrapper/pkg/slack"
	"time"
)

// Configure a init settings as env variable and other for launch worker
func initConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	return nil
}

//Building Message for send to slack
func buildMessage(covidInfo []entity.Covid19InfoEntity) (msg slackUtil.MessageAttachmentsFormat) {
	msg = slackUtil.MessageAttachmentsFormat{
		Color:   "#36a64f",
		Text:    fmt.Sprint("오늘까지의 누적 확진자 수는 ", covidInfo[0].DecideCnt, "명 입니다. :sob:"),
	}
	return
}

func dbInitConfig() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	// Migrate the schema
	_ = db.AutoMigrate(&entity.Covid19InfoEntity{})
	return db
}

func main() {
	fmt.Println("Start")
	_ = initConfig()
	db := dbInitConfig()

	var covidInfoArr []entity.Covid19InfoEntity
	startDt := time.Now().AddDate(0,0, -1) // yesterday
	endDt := time.Now()
	covidInfoArr = scrape.Scrape(os.Getenv("OPEN_API_KEY"), startDt, endDt)
	//todayCovidInfo := scrape.MakeMockCovidInfo()

	db.Create(&covidInfoArr)

	fmt.Println(covidInfoArr)
	fmt.Println("End")
	//AlarmToSlack()
}

/*
func AlarmToSlack() {
	todayCovidInfo := scrape.Scrape(os.Getenv("OPEN_API_KEY"))
	builtMessage := buildMessage(todayCovidInfo)
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	_ = slackUtil.SendSlackMessage(slackWebhookUrl,"bot-test", "오늘의 코로나 알림 :mask:", []slackUtil.MessageAttachmentsFormat{builtMessage})
}
 */