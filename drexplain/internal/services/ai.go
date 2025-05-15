package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetChatCompletions(prompt string, apiKey string) (string, error) {
	url := "https://openrouter.ai/api/v1/chat/completions"

	body := map[string]any{
		"model": "deepseek/deepseek-r1:free",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Choices[0].Message.Content, nil
}
