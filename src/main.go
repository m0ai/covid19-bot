package main

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	entity "scrapper/internal/entity"
	dbcontext "scrapper/pkg/dbcontext"
	slack "scrapper/pkg/slack"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Notify Start")
	db := dbcontext.DbInitConfig()
	AlarmToSlack(strconv.Itoa(getTodayDecideCnt(db)))
	fmt.Println("Notify End")
}

// getTodayDecideCnt
func getTodayDecideCnt(db *gorm.DB) int {
	today := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1) // yesterday

	var todayEntity entity.Covid19InfoEntity
	var yesterdayEntity entity.Covid19InfoEntity
	db.Order("state_dt desc").First(&todayEntity, "state_dt <= ?", today)
	db.Order("state_dt desc").First(&yesterdayEntity, "state_dt <= ?", yesterday)
	return todayEntity.DecideCnt - yesterdayEntity.DecideCnt
}

func AlarmToSlack(msg string) {
	builtMessage := buildSlackMessage(msg)
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	_ = slack.SendSlackMessage(slackWebhookUrl, "bot-test", "오늘의 코로나 알림 :mask:", []slack.MessageAttachmentsFormat{builtMessage})
}

// Building Message for send to slack
func buildSlackMessage(msg string) (slackMsg slack.MessageAttachmentsFormat) {
	slackMsg = slack.MessageAttachmentsFormat{
		Color:  "#36a64f",
		Text:   fmt.Sprint("금일 신규 확진자 수는 ", msg, "명 입니다. :sob:"),
		Footer: "Data From Open API",
	}
	return
}
