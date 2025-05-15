package handlers

import (
	"drexplain/internal/services"
	"drexplain/internal/utils"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDefine(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	query := msg.CommandArguments()
	apiKey, _ := utils.GetEnvString("R1_KEY", "")
	if query == "" {
		utils.Reply(bot, msg, "Usage: /define [term]")
		return
	}

	utils.Reply(bot, msg, "Definition of "+query+" goes here.")
	prompt := fmt.Sprintf("Define the term between >>><<<, in a short, clear and coincise way: >>>%s<<<", query)
	response, err := services.GetChatCompletions(prompt, apiKey)
	if err != nil {
		fmt.Println("LLM error: ", err.Error())
		utils.Reply(bot, msg, "Cannot connect to LLM")
		return
	}

	utils.Reply(bot, msg, response)
}
