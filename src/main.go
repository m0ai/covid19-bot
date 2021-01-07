package main

import (
	"fmt"
	"os"
	entity "scrapper/internal/entity"
	dbcontext "scrapper/pkg/dbcontext"
	slack "scrapper/pkg/slack"
	"time"
)

func main() {
	fmt.Println("Notify Start")

	todayEntity, _ := getCovid19infoEntityFilterByStateDt(time.Now())
	yesterdayEntity, _ := getCovid19infoEntityFilterByStateDt(time.Now().AddDate(0, 0, -1))

	builtMessage := buildSlackMessage(
		todayEntity.DecideCnt-yesterdayEntity.DecideCnt,
		todayEntity.DeathCnt-yesterdayEntity.DeathCnt,
	)

	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	AlarmToSlack("general", []slack.MessageAttachmentsFormat{builtMessage}, slackWebhookUrl)
	fmt.Println("Notify End")
}

func getCovid19infoEntityFilterByStateDt(dt time.Time) (entity.Covid19InfoEntity, error) {
	db := dbcontext.DbInitConfig()
	var covid19info entity.Covid19InfoEntity
	db.Order("state_dt desc").First(&covid19info, "state_dt <= ?", dt)
	return covid19info, nil
}

func AlarmToSlack(channel string, attachmentMessages []slack.MessageAttachmentsFormat, slackWebhookUrl string) {
	_ = slack.SendSlackMessage(slackWebhookUrl, "bot-test", "오늘의 코로나 알림 :mask:", attachmentMessages)
}

// Building Message for send to slack
func buildSlackMessage(decideCnt, DeathCnt int) (slackMsg slack.MessageAttachmentsFormat) {
	slackMsg = slack.MessageAttachmentsFormat{
		Color:      "#36a64f",
		Title:      "국내 코로나 19 바이러스 일별 정보",
		TitleLink:  "http://ncov.mohw.go.kr",
		Footer:     "공공 데이터 포털 제공",
		FooterLink: "https://www.data.go.kr",
		Ts:         time.Now().Unix(),
		Fields: []slack.InnerField{
			{Title: "신규 확진자 수", Value: fmt.Sprintln(decideCnt, "명"), Short: false},
			{Title: "신규 사망자 수", Value: fmt.Sprintln(DeathCnt, "명"), Short: false},
		},
	}
	return
}
