package services

import (
	"chronobot/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type NasaData struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func FetchNasaPhoto(date string) (NasaData, error) {
	nasaKey := utils.GetEnvString("NASA_KEY", "")
	url := fmt.Sprintf("https://api.nasa.gov/planetary/apod?api_key=%s&date=%s", nasaKey, date)
	resp, err := http.Get(url)
	if err != nil {
		return NasaData{Title: "Not available", URL: ""}, fmt.Errorf("cannot access NASA API: %v", err)
	}
	defer resp.Body.Close()
	var data NasaData
	json.NewDecoder(resp.Body).Decode(&data)

	return data, nil
}
