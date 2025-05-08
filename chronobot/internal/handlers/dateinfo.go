package handlers

import (
	"chronobot/internal/services"
	"fmt"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleDateInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update, t time.Time) {
	body := buildResponse(t)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, body)
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func buildResponse(t time.Time) string {
	now := time.Now()

	daysAgo := int(now.Sub(t).Hours() / 24)
	weekday := t.Weekday().String()

	zodiac := services.GetZodiac(t)
	chineseZodiac := services.GetChineseZodiac(t.Year())

	if daysAgo < 0 {
		// Calculate days until the date
		daysUntil := int(t.Sub(now).Hours() / 24)
		return fmt.Sprintf("🕰️ *%s* — that's in %d days!\n\n📆 It will be a *%s*\n♈ People born on this day are *%s*\n🐲 In Chinese zodiac, they'd be a *%s*",
			t.Format("2006-01-02"), daysUntil, weekday, zodiac, chineseZodiac)
	}

	event, birthdays, err := services.FetchOnThisDay(t.Month(), t.Day())
	eventText := fmt.Sprintf(" 📜 Back in time, here's recent event happend on %s %2d:\n %s\n", t.Month().String(), t.Day(), event)
	birthdayText := fmt.Sprintf(" 🎂 Famous birthdays on this day:\n %s\n", birthdays)
	if err != nil {
		fmt.Println("Error fetching events:", err)
		eventText = ""
		birthdayText = ""
	}

	dateInfo := fmt.Sprintf(` 🕰️ *%s* - that's %d days ago!
  📅 It was a *%s*
  ♈ People born on this day are *%s*
  🐲 In Chinese zodiac, they'd be a *%s*
  `,
		t.Format("2006-01-02"), daysAgo, weekday, zodiac, chineseZodiac)

	if eventText != "" {
		dateInfo += eventText
	}
	if birthdayText != "" {
		dateInfo += birthdayText
	}

	nasa, err := services.FetchNasaPhoto(t.Format("2006-01-02"))
	if err == nil {
		dateInfo += fmt.Sprintf("🌌 NASA's Picture of the Day: %s\n", nasa)
	} else {
		fmt.Println("Error fetching NASA photo:", err)
	}

	return dateInfo
}
