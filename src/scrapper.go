package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	entity "scrapper/internal/entity"
	dbcontext "scrapper/pkg/dbcontext"
	scrape "scrapper/pkg/scrapper"
	"time"
)

func main() {
	fmt.Println("Scrapper Start")
	db := dbcontext.DbInitConfig()
	startDt := time.Now().AddDate(0, 0, -3) // three days ago
	endDt := time.Now()
	covid19InfoArr := scrape.Scrape(os.Getenv("OPEN_API_KEY"), startDt, endDt)
	upsertToDB(db, covid19InfoArr)
	fmt.Println("Scrapper End")
}

func upsertToDB(db *gorm.DB, covid19infoArr []entity.Covid19InfoEntity) {
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&covid19infoArr)
}
