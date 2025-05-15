package utils

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Reply(bot *tgbotapi.BotAPI, message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	bot.Send(msg)
}
