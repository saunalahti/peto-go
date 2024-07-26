package main

import (
	"context"
	"encoding/json"
	"time"
)

func main() {
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
