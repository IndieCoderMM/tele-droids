package bot

import (
	"chronobot/internal/utils"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func InitBot() *tgbotapi.BotAPI {
	token := utils.GetEnvString("TG_BOT_KEY", "")
	if token == "" {
		log.Fatal("Telegram bot token is not set.")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account: %s", bot.Self.UserName)

	return bot
}
