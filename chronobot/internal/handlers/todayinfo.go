package handlers

import (
	"chronobot/internal/services"
	"chronobot/internal/utils"
	"fmt"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleTodayInfo(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	t := time.Now()
	zodiac := services.GetZodiac(t)
	chineseZodiac := services.GetChineseZodiac(t.Year())

	body := fmt.Sprintf("📅 *%s* - *%s*\n"+
		"♈ People born on this day are *%s*\n"+
		"🐲 In Chinese zodiac, they'd be a *%s*\n",
		t.Format("2006-01-02"), t.Weekday().String(), zodiac, chineseZodiac)

	milestones := utils.DaysUntil(t)

	body += fmt.Sprintf("\n🗓️ Days until:\n"+
		"🌙 Next month: *%d days*\n"+
		"☀️ Next year: *%d days*\n"+
		"🔟 Next decade: *%d days*\n"+
		"🕰️ Next century: *%d days*\n"+
		"🛸 Next millennium: *%d days*\n",
		milestones.DaysToNextMonth,
		milestones.DaysToNextYear,
		milestones.DaysToNextDecade,
		milestones.DaysToNextCentury,
		milestones.DaysToNextMillennium,
	)

	birthdays := services.FetchBirthdays(t.Month(), t.Day())
	if birthdays != "" {
		body += fmt.Sprintf("\n🎂 Famous birthdays on this day:\n%s\n", birthdays)
	}

	events := services.FetchEvent(t.Month(), t.Day())
	if events != "" {
		body += fmt.Sprintf("\n📜 Today in history:\n %s\n", events)
	}

	nasa, err := services.FetchNasaPhoto(t.Format("2006-01-02"))
	if err == nil && nasa.URL != "" {
		body += fmt.Sprintf("\n🌌 NASA's Picture of the Day:\n[%s](%s)\n", nasa.Title, nasa.URL)
	} else {
		fmt.Println("Error fetching NASA photo:", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, body)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}
