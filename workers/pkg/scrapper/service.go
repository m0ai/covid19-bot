package scrapper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"scrapper/internal/entity"
	"time"
)

type Response struct {
	XMLName xml.Name 				 `xml:"response"`
	Items []entity.Covid19InfoEntity `xml:"body>items>item"`
}

// Scrapping Covid-19 Information from OPEN API
// https://www.thepolyglotdeveloper.com/2017/03/parse-xml-data-in-a-golang-application/
func Scrape(openAPIKey string, startDt, endDt time.Time) []entity.Covid19InfoEntity {
	rawData := getCovidDataFromAPI(openAPIKey, startDt, endDt)
	extractedCovid19Data := extractData(rawData)
	return extractedCovid19Data
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
func getCovidDataFromAPI(openAPIKey string, startDate, endDate time.Time) []byte {
	var baseUrl string = "http://openapi.data.go.kr/openapi/service/rest/Covid19/getCovid19InfStateJson"
	params := map[string]string{
		"ServiceKey":    openAPIKey,
		"startCreateDt": startDate.Format("20060102"),
		"endCreatDt":    endDate.Format("20060102"),
		"pageNo":        "1",
		"numOfRows":     "10",
	}
	xmlData, _ := requestTo(baseUrl, params)
	return xmlData
}

func extractData (data []byte) []entity.Covid19InfoEntity {
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

func extractFirstData (data []byte) entity.Covid19InfoEntity {
	return extractData(data)[0]
}

func MakeMockCovid19Data() []entity.Covid19InfoEntity {
	data := entity.Covid19InfoEntity {
		Seq: 33,
		DecideCnt: 1,
		DeathCnt: 1,
	}

	return []entity.Covid19InfoEntity{data}
}