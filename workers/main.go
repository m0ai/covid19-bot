package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"os"
	"scrapper/internal/entity"
	scrape "scrapper/pkg/scrapper"
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
	dsn := fmt.Sprint("" +
		"host=postgres", " ",
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
	_ = initConfig()
	db := dbInitConfig()

	startDt := time.Now().AddDate(0,0, -1) // yesterday
	endDt := time.Now()
	covid19InfoArr := scrape.Scrape(os.Getenv("OPEN_API_KEY"), startDt, endDt)
	// covid19InfoArr := scrape.MakeMockCovid19Data()

	upsertToDB(db, covid19InfoArr)
	//AlarmToSlack()
	fmt.Println("End")
}

func upsertToDB(db *gorm.DB, covid19infoArr []entity.Covid19InfoEntity) {
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&covid19infoArr)
}

/*
func AlarmToSlack() {
	todayCovidInfo := scrape.Scrape(os.Getenv("OPEN_API_KEY"))
	builtMessage := buildMessage(todayCovidInfo)
	slackWebhookUrl := os.Getenv("SLACK_WEBHOOK_URL")
	_ = slackUtil.SendSlackMessage(slackWebhookUrl,"bot-test", "오늘의 코로나 알림 :mask:", []slackUtil.MessageAttachmentsFormat{builtMessage})
}

//Building Message for send to slack
func buildMessage(covidInfo []entity.Covid19InfoEntity) (msg slackUtil.MessageAttachmentsFormat) {
	msg = slackUtil.MessageAttachmentsFormat{
		Color:   "#36a64f",
		Text:    fmt.Sprint("오늘까지의 누적 확진자 수는 ", covidInfo[0].DecideCnt, "명 입니다. :sob:"),
	}
	return
}
 */