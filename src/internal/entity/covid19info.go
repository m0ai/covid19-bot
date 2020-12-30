package entity

import (
	"database/sql/driver"
	"encoding/xml"
	"gorm.io/gorm"
	"time"
)

type timeAsKST struct {
	time.Time
}

type stateDt struct {
	time.Time
}

type stateTime struct {
	time.Time
}

type Covid19InfoEntity struct {
	gorm.Model
	XMLName   xml.Name `xml:"item" json:"Item" gorm:"-"`
	Seq       int      `xml:"seq" gorm:"primaryKey; uniqueIndex; not null"`
	DecideCnt int      `xml:"decideCnt"` // 누적 확진자 수
	DeathCnt  int      `xml:"deathCnt"`  // 사망자 수
	CareCnt   int      `xml:"careCnt"`   // 치료중 환자 수
	ClearCnt  int      `xml:"clearCnt"`  // 격리 해제 수

	CreateDt  timeAsKST `xml:"createDt" gorm:"type:timestamp"` // 2020-12-22 09:35:08.23
	StateDt   stateDt   `xml:"stateDt" gorm:"type:date"`       // 2020-12-22
	StateTime stateTime `xml:"stateTime" gorm:"type:time"`     // 00:00

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity Covid19InfoEntity) TableName() string {
	return "covid19info"
}

// UnmarshalXML a StateDt, StateDt Does have only 'YearMonthDay' Field (e.g "2020-12-22")
func (s *stateDt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var v string
	loc, _ := time.LoadLocation("Asia/Seoul")
	_ = d.DecodeElement(&v, &start)
	parse, _ := time.ParseInLocation("20060102", v, loc)
	*s = stateDt{parse}
	return nil
}

func (s *stateDt) Scan(src interface{}) error {
	if _time, ok := src.(time.Time); ok {
		s.Time = _time
	}
	return nil
}

func (s stateDt) Value() (driver.Value, error) {
	return s.Time, nil
}

// UnmarshalXML with a Doesn't have 'T' (as Timezone) datetime Field (e.g "2020-12-22 00:00:00.23")
func (t *timeAsKST) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	_ = d.DecodeElement(&v, &start)
	parse, _ := time.Parse("2006-01-02 15:04:05.00", v)
	*t = timeAsKST{parse}
	return nil
}

func (t *timeAsKST) Scan(src interface{}) error {
	if _time, ok := src.(time.Time); ok {
		t.Time = _time
	}
	return nil
}

func (t timeAsKST) Value() (driver.Value, error) {
	return t.Time, nil
}

func (s *stateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	loc, _ := time.LoadLocation("Asia/Seoul")
	_ = d.DecodeElement(&v, &start)
	parse, _ := time.ParseInLocation("15:04", v, loc)
	*s = stateTime{parse}
	return nil
}

func (s *stateTime) Scan(src interface{}) error {
	if _time, ok := src.(time.Time); ok {
		s.Time = _time
	}
	return nil
}

func (s stateTime) Value() (driver.Value, error) {
	return s.Time, nil
}
