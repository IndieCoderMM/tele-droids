package bot

import (
	"chronobot/internal/handlers"
	"chronobot/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetWebhook(bot *tgbotapi.BotAPI, url string) error {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		return fmt.Errorf("failed to set webhook: %v", err)
	}
	return nil
}

func WebhookHandler(bot *tgbotapi.BotAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var update tgbotapi.Update
		// Parse incoming JSON payload (Telegram updates)
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			log.Println("Error decoding update:", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if update.Message == nil {
			return
		}

		if update.Message.IsCommand() {
			handleCommand(bot, update)
		} else {
			t, err := utils.ParseDate(update.Message.Text)
			if err != nil {
				handlers.HandleHelp(bot, update)
			}

			// Handle date input
			handlers.HandleDateInfo(bot, update, t)
		}

		// Respond to Telegram
		w.WriteHeader(http.StatusOK)
	}
}

func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		handlers.HandleStart(bot, update)
	case "help":
		handlers.HandleHelp(bot, update)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Please send me a date in YYYY-MM-DD format.")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		bot.Send(msg)
	}
}
