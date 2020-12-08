package scrapper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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
	xmlFp, err := os.Open(xmlFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Sprintln("Successfully Opened a ", xmlFilePath)
	defer xmlFp.Close()

	var res Response
	data, err := ioutil.ReadAll(xmlFp)
	xmlReadErr := xml.Unmarshal(data, &res)
	if xmlReadErr != nil {
		panic(xmlReadErr)
	}

	for _, item := range res.Items {
		fmt.Println(item)
	}
}