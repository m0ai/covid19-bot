package main

import (
	"fmt"
	"scrapper/scrapper"
)

func main() {
	fmt.Println("Start")
	scrapper.Scrape("dump.xml")
	fmt.Println("End")
}
