package main

import (
	b "drexplain/internal/bot"
	"drexplain/internal/handlers"
	"drexplain/internal/utils"
	"encoding/json"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token, ok := utils.GetEnvString("TG_BOT_KEY", "")
	if !ok {
		log.Panic("Bot token not found")
	}

	url, ok := utils.GetEnvString("DOMAIN_URL", "")
	if !ok {
		log.Panic("Bot token not found")
	}

	bot := b.InitBot(token)
	if err := b.InitWebhook(bot, url+"/webhook"); err != nil {
		log.Panic(err)
	}

	// healthcheck
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var update tgbotapi.Update
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			log.Println("Error decoding update:", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if update.Message == nil || !update.Message.IsCommand() {
			return
		}

		handleCommand(bot, update.Message)
		w.WriteHeader(http.StatusOK)
	},
	)

	port, _ := utils.GetEnvString("PORT", "8080")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "define":
		handlers.HandleDefine(bot, message)
	case "start":
		utils.Reply(bot, message, "Welcome to Drexplain! How can I assist you today?")
	case "help":
		utils.Reply(bot, message, "Available commands:\n[/start](start) - Start the bot\n[/help](help) - Show available commands")
	default:
		utils.Reply(bot, message, "Unknown command. Type [/help](help) for a list of available commands.")
	}
}
