package main

import (
	"chronobot/internal/bot"
	"chronobot/internal/utils"
	"log"
	"net/http"
)

func main() {
	port := utils.GetEnvString("PORT", ":8080")
	url := utils.GetEnvString("DOMAIN_URL", "")

	b := bot.InitBot()

	if err := bot.SetWebhook(b, url+"/webhook"); err != nil {
		log.Fatal("Failed to set webhook:", err)
	}

	http.HandleFunc("/webhook", bot.WebhookHandler(b))
	log.Fatal(http.ListenAndServe(port, nil))
}
