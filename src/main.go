package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"os"
	entity "scrapper/internal/entity"
	scrape "scrapper/pkg/scrapper"
	slack "scrapper/pkg/slack"
	"strconv"
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

func dbInitConfig() *gorm.DB {
	dsn := fmt.Sprint(""+
		"host=", os.Getenv("POSTGRES_HOST"), " ",
		"dbname=", os.Getenv("POSTGRES_DB"), " ",
		"user=", os.Getenv("POSTGRES_USER"), " ",
		"password=", os.Getenv("POSTGRES_PASSWORD"), " ",
		"port=5432 sslmode=disable TimeZone=Asia/Seoul connect_timeout=15")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Covid19InfoEntity{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func main() {
	fmt.Println("Start")
	// _ = initConfig()
	db := dbInitConfig()

	startDt := time.Now().AddDate(0, 0, -3) // yesterday
	endDt := time.Now()
	covid19InfoArr := scrape.Scrape(os.Getenv("OPEN_API_KEY"), startDt, endDt)

	fmt.Println(covid19InfoArr)

	upsertToDB(db, covid19InfoArr)

	AlarmToSlack(strconv.Itoa(getTodayDecideCnt(db)))
	fmt.Println("End")
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

func upsertToDB(db *gorm.DB, covid19infoArr []entity.Covid19InfoEntity) {
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&covid19infoArr)
}

func AlarmToSlack(msg string) {
	builtMessage := buildSlackMessage(msg)
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	_ = slack.SendSlackMessage(slackWebhookUrl, "bot-test", "오늘의 코로나 알림 :mask:", []slack.MessageAttachmentsFormat{builtMessage})
}

// Building Message for send to slack
func buildSlackMessage(msg string) (slackMsg slack.MessageAttachmentsFormat) {
	slackMsg = slack.MessageAttachmentsFormat{
		Color: "#36a64f",
		Text:  fmt.Sprint("신규 확진자 수는 ", msg, "명 입니다. :sob:"),
	}
	return
}
