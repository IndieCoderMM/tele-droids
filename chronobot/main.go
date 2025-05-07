package main

import (
	"chronobot/utils"
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"time"
)

func main() {
	port := utils.GetEnvString("PORT", ":8080")
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

	url := utils.GetEnvString("DOMAIN_URL", "")
	err = setWebhook(bot, url)
	if err != nil {
		log.Fatal("Failed to set webhook:", err)
	}

	http.HandleFunc("/hook", handler(bot))
	log.Fatal(http.ListenAndServe(port, nil))
}

func setWebhook(bot *tgbotapi.BotAPI, url string) error {
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(url))
	if err != nil {
		return fmt.Errorf("failed to set webhook: %v", err)
	}
	return nil
}

func handler(bot *tgbotapi.BotAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse incoming JSON payload (Telegram updates)
		update := tgbotapi.Update{}
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			log.Println("Error decoding update:", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Handle new user messages
		if update.Message != nil {
			// Check if the message is a command
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to ChronoBot! Please send me a date in YYYY-MM-DD format.")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msg)
				case "help":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I can provide you with historical events and zodiac signs for a given date. Just send me a date in YYYY-MM-DD format.")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msg)
				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command. Please send me a date in YYYY-MM-DD format.")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msg)
				}
			} else {
				// Regular date input handling
				input := update.Message.Text
				t, err := utils.ParseDate(input)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid date format. Please use YYYY-MM-DD.")
					bot.Send(msg)
					return
				}

				// Send the response
				response := buildResponse(t)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
				msg.ParseMode = "Markdown"
				bot.Send(msg)
			}
		}

		// Respond to Telegram
		w.WriteHeader(http.StatusOK)
	}
}

func buildResponse(t time.Time) string {
	now := time.Now()

	daysAgo := int(now.Sub(t).Hours() / 24)
	weekday := t.Weekday().String()
	zodiac := utils.GetZodiac(t)
	chineseZodiac := utils.GetChineseZodiac(t.Year())
	onThisDay := utils.FetchOnThisDay(t.Month(), t.Day())
	nasa := utils.FetchNasaPhoto(t.Format("2006-01-02"))

	return fmt.Sprintf("📅 *%s* — %d days ago\n\n🗓️ Day of the week: *%s*\n♒ Western Zodiac: *%s*\n🐉 Chinese Zodiac: *%s*\n\n🧠 *On This Day:* %s\n\n📷 *NASA Photo:* %s — %s",
		t.Format("2006-01-02"), daysAgo, weekday, zodiac, chineseZodiac, onThisDay, nasa.Title, nasa.URL)
}
