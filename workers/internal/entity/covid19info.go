package entity

import (
	"encoding/xml"
	"gorm.io/gorm"
	"time"
)

type Covid19InfoEntity struct {
	gorm.Model
	XMLName   xml.Name `xml:"item" json:"Item" gorm:"-"`
	Seq       int `xml:"seq" gorm:"index"`
	DecideCnt int `xml:"decideCnt"` // 누적 확진자 수
	DeathCnt  int `xml:"deathCnt"` // 사망자 수
	CareCnt   int `xml:"careCnt"` // 치료중 환자 수
	ClearCnt  int `xml:"clearCnt"` // 격리 해제 수
	StateDt   string `xml:"stateDt"`
	StateTime string `xml:"stateTime"`
	CreateDt  string `xml:"createDt"`

	CreatedAt time.Time
	UpdatedAt time.Time
	TodayDecideCnt int
}

