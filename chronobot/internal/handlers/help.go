package handlers

import (
	"chronobot/internal/templates"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, templates.Help)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}
