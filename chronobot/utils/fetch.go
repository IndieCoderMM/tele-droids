package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Fetch an event from Wikipedia's On This Day API
func FetchOnThisDay(month time.Month, day int) string {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/events/%2d/%2d", month, day)

	type Event struct {
		Text string `json:"text"`
	}
	var result struct {
		Events []Event `json:"events"`
	}

	resp, err := http.Get(url)
	if err != nil {
		return "Cannot access wikipedia API"
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	if len(result.Events) == 0 {
		return "No events found for this date."
	}

	return result.Events[0].Text
}

type NasaData struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func FetchNasaPhoto(date string) NasaData {
	nasaKey := GetEnvString("NASA_KEY", "")
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&date=%s", nasaKey, date)
	resp, err := http.Get(url)
	if err != nil {
		return NasaData{Title: "Not available", URL: ""}
	}
	defer resp.Body.Close()
	var data NasaData
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}
