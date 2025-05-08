package handlers

import (
	"chronobot/internal/services"
	"fmt"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleDateInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update, t time.Time) {
	body := buildResponse(t)
	fmt.Println("Response body", body)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, body)
	msg.ParseMode = "Markdown"
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

	dateInfo := fmt.Sprintf("🕰️ *%s* - that's %d days ago!\n\n"+
		"- 📅 It was a *%s*\n"+
		"- ♈ People born on this day are *%s*\n"+
		"- 🐲 In Chinese zodiac, they'd be a *%s*\n",
		t.Format("2006-01-02"), daysAgo, weekday, zodiac, chineseZodiac)

	birthdays := services.FetchBirthdays(t.Month(), t.Day())
	if birthdays != "" {
		fmt.Println("Birthdays", birthdays)
		dateInfo += fmt.Sprintf("\n🎂 Famous birthdays on this day:\n%s\n", birthdays)
	}

	events := services.FetchEvent(t.Month(), t.Day())
	if events != "" {
		fmt.Println("Events", birthdays)
		dateInfo += fmt.Sprintf("\n📜 Back in time:\n %s\n", events)
	}

	nasa, err := services.FetchNasaPhoto(t.Format("2006-01-02"))
	if err == nil {
		dateInfo += fmt.Sprintf("\n🌌 NASA's Picture of the Day:\n[%s](%s)\n", nasa.Title, nasa.URL)
	} else {
		fmt.Println("Error fetching NASA photo:", err)
	}

	return dateInfo
}
