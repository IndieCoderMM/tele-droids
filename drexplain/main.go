package main

import (
	b "drexplain/internal/bot"
	"drexplain/internal/handlers"
	"drexplain/internal/utils"
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

	port, _ := utils.GetEnvString("PORT", ":8080")

	bot := b.InitBot(token)
	if err := b.InitWebhook(bot, url); err != nil {
		log.Panic(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	log.Fatal(http.ListenAndServe(port, nil))

	for update := range updates {
		if update.Message == nil || !update.Message.IsCommand() {
			continue
		}

		handleCommand(bot, update.Message)
	}
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
