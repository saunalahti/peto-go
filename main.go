package main

import (
	"context"
	"encoding/json"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	validateEnv()

	ConnectCache()
	ScrapeAndSave()

	go StartWeb()

	for range time.Tick(30 * time.Second) {
		ScrapeAndSave()
	}
}

func ScrapeAndSave() {
	data, err := Scrape()

	if err != nil {
		println("failed to scrape")
	}

	cacheData, _ := Cache.Get(context.Background(), "peto-data")

	marshalData, _ := json.Marshal(data)

	if cacheData != string(marshalData) {
		println("New content, updating.")
		Cache.Set(context.Background(), "peto-data", string(marshalData))
	}
}

func validateEnv() {
	godotenv.Load(".env")

	url := os.Getenv("SCRAPE_URL")

	if url != "" {
		regex := regexp.MustCompile(`^https?:\/\/[^\s/$.?#].[^\s]*$`)

		if !regex.MatchString(url) {
			panic("variable SCRAPE_URL is set, but the URL is invalid. It must start with http(s).")
		}
	}
}
