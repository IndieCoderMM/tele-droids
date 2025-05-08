package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Fetch from Wikipedia's On This Day API
func FetchOnThisDay(month time.Month, day int) (string, string, error) {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/en/onthisday/all/%2d/%2d", month, day)

	type Event struct {
		Text string `json:"text"`
		Year string `json:"year"`
	}
	var result struct {
		Events []Event `json:"events"`
		Births []Event `json:"births"`
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("cannot access Wikipedia server: %v", err)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)

	var eventText, birthText string

	if len(result.Events) == 0 {
		eventText = "No events found for this date."
	} else {
		eventText = fmt.Sprintf("%s (%s)", result.Events[0].Text, result.Events[0].Year)
	}

	if len(result.Births) == 0 {
		birthText = "No births found for this date."
	} else {
		for i := 0; i < len(result.Births) && i < 3; i++ {
			birthText = fmt.Sprintf("- %s\n", result.Births[i].Text)
		}
	}

	return eventText, birthText, nil
}
