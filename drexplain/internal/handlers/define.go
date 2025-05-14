package handlers

import (
	"drexplain/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDefine(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	query := msg.CommandArguments()
	if query == "" {
		utils.Reply(bot, msg, "Usage: /define [term]")
		return
	}

	// TODO: Call LLM
	utils.Reply(bot, msg, "Definition of "+query+" goes here.")
}
