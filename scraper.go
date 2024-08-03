package main

import (
	"fmt"
	"peto/models"
	"peto/util"
	"strings"

	"github.com/gocolly/colly/v2"
)

func splitByDelimiters(r rune) bool {
	return r == ',' || r == ':'
}

func Scrape() ([]models.Event, error) {
	var events []models.Event

	var respError error

	c := colly.NewCollector()
	c.DetectCharset = true

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		data := strings.Split(e.Text, "\n\n")

		for _, v := range data {
			lines := strings.Split(strings.TrimSpace(v), "\n")

			if len(lines) >= 3 {
				locations := strings.Split(lines[0], "/")

				incidentParts := strings.FieldsFunc(lines[1], splitByDelimiters)

				event := models.Event{
					Location: locations[0],
					Incident: strings.TrimSpace(incidentParts[0]),
					Datetime: lines[2],
				}

				if locations[0] != locations[1] {
					event.LocationInt = locations[1]
				}

				if len(incidentParts) == 2 {
					severity := strings.TrimSpace(incidentParts[1])

					switch severity {
					case "pieni":
						event.Severity = 3
					case "keskisuuri":
						event.Severity = 2
					case "suuri":
						event.Severity = 1
					}
				}

				event.ID = util.GenerateID(event.Location, event.Datetime)

				events = append(events, event)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)

		respError = err
	})

	c.Visit("https://www.peto-media.fi/")

	if respError != nil {
		return []models.Event{}, respError
	}

	return events, nil
}
