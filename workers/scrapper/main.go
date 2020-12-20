package scrapper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Item struct {
	XMLName   xml.Name `xml:"item" json:"Item" db:""`
	Seq       int `xml:"seq"`
	DecideCnt int `xml:"decideCnt"` // 누적 확진자 수
	DeathCnt  int `xml:"deathCnt"` // 사망자 수
	CareCnt   int `xml:"careCnt"` // 치료중 환자 수
	ClearCnt  int `xml:"clearCnt"` // 격리 해제 수
	StateDt   string `xml:"stateDt"`
	StateTime string `xml:"stateTime"`
	CreateDt  string `xml:"createDt"`

	TodayDecideCnt int
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Items []Item `xml:"body>items>item"`
}

// Scrapping a Covid-19 Information
// https://www.thepolyglotdeveloper.com/2017/03/parse-xml-data-in-a-golang-application/
func Scrape(openAPIKey string) Item {
	startDt := time.Now().AddDate(0,0, -1)
	endDt := time.Now()
	item := getCovidDataFromAPI(openAPIKey, startDt, endDt)
	// item := makeCovid19MockStruct()
	return item
}

func makeCovid19MockStruct() *Item {
	item := Item{
		Seq:       1,
		CareCnt:   3,
		DeathCnt:  2,
		DecideCnt: 2,
	}
	return &item
}

func requestTo(baseUrl string, params map[string]string) ([]byte, error) {
	i := 0
	for k, v := range params {
		if i == 0 {
			baseUrl = fmt.Sprint(baseUrl, "?")
		} else {
			baseUrl = fmt.Sprint(baseUrl, "&")
		}

		baseUrl = fmt.Sprint(baseUrl, k, "=", v)
		i++
	}
	res, err := http.Get(baseUrl)

	fmt.Println("Request to " + baseUrl)
	if err != nil {
		log.Fatalln("GET Request Failed GET (" + baseUrl + ")")
	}
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with status: ", res.StatusCode)
	}

	resBody, readErr := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if readErr != nil {
		log.Fatalln("Response data Error")
	}
	return resBody, nil
}

// curl --include --request GET '
//http://openapi.data.go.kr/openapi/service/rest/Covid19/getCovid19InfStateJson?
// ServiceKey=${}
// pageNo=1
// numOfRows=10
// startCreateDt=20200310
// endCreateDt=20200315
func getCovidDataFromAPI(openAPIKey string, startDate, endDate time.Time) (extractedCovid19Data Item){
	var baseUrl string = "http://openapi.data.go.kr/openapi/service/rest/Covid19/getCovid19InfStateJson"
	params := map[string]string{
		"ServiceKey":    openAPIKey,
		"startCreateDt": startDate.Format("20060102"),
		"endCreatDt":    endDate.Format("20060102"),
		"pageNo":        "1",
		"numOfRows":     "10",
	}
	xmlData, _ := requestTo(baseUrl, params)
	extractedCovid19Data = extractFirstData(xmlData)
	return
}

func extractData (data []byte) []Item {
	var res Response
	xmlReadErr := xml.Unmarshal(data, &res)
	if xmlReadErr != nil {
		panic(xmlReadErr)
	}

	for _, item := range res.Items {
		fmt.Println(item)
	}
	return res.Items
}

func extractFirstData (data []byte) Item {
	return extractData(data)[0]
}

