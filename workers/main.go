package main

import (
	"fmt"
	"log"
	"os"
	"scrapper/scrapper"
	"github.com/joho/godotenv"
)

func getEnvVariableFromFile(key, envFile string) string {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	apiKey := getEnvVariableFromFile("OPEN_API_KEY", "../.env")
	fmt.Println(apiKey)
	fmt.Println("Start")
	scrapper.Scrape("dump.xml")
	fmt.Println("End")
}
