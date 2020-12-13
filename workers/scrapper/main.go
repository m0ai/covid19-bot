package scrapper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Item []struct {
	XMLName   xml.Name `xml:"item"`
	Seq       int `xml:"seq"`
	DecideCnt int `xml:"decideCnt"`
	DeathCnt  int `xml:"deathCnt"`
	CareCnt   int `xml:"careCnt"`
	StateDt   string `xml:"stateDt"`
	StateTime string `xml:"stateTime"`
	CreateDt  string `xml:"createDt"`
}

// Xml => response != Response
type Response struct {
	XMLName xml.Name `xml:"response"`
	Items []Item `xml:"body>items>item"`
}


// Scrapping a Covid-19 Information
// https://www.thepolyglotdeveloper.com/2017/03/parse-xml-data-in-a-golang-application/
func Scrape(xmlFilePath string) {
	t := time.Now()
	getCovidDataFromAPI(t, t)
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
func getCovidDataFromAPI(startDate, endDate time.Time) {
	var baseUrl string = "http://openapi.data.go.kr/openapi/service/rest/Covid19/getCovid19InfStateJson"
	params := map[string]string{
		"ServiceKey" 	: os.Getenv("OPEN_API_KEY"),
		"startCreateDt" : startDate.Format("20060102"),
		"endCreatDt" 	: endDate.Format("20060102"),
		"pageNo"	    : "1",
		"numOfRows"		: "10",
	}

	xmlData, _ := requestTo(baseUrl, params)
	_ = extractData(xmlData)
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

