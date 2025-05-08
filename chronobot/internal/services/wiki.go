package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Fetch from Wikipedia's On This Day API
func FetchBirthdays(month time.Month, day int) string {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/births/%2d/%2d", month, day)

	type Event struct {
		Text string `json:"text"`
	}

	var result struct {
		Births []Event `json:"births"`
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return ""
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)

	var birthdays string

	if len(result.Births) == 0 {
		birthdays = ""
	} else {
		for i := 0; i < len(result.Births) && i < 3; i++ {
			birthdays += fmt.Sprintf("%d. %s\n", i+1, result.Births[i].Text)
		}
	}

	return birthdays
}

func FetchEvent(month time.Month, day int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/%d/date", month, day)
	// Response is plain text with no body
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ""
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(bodyBytes)
}
