package main

import (
	"chronobot/utils"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
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

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message Updates
			continue
		}

		input := update.Message.Text
		t, err := utils.ParseDate(input)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid date format. Please use YYYY-MM-DD.")
			bot.Send(msg)
			continue
		}

		response := buildResponse(t)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ParseMode = "Markdown"
		bot.Send(msg)
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
