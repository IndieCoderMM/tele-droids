package handlers

import (
	"chronobot/internal/services"
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

	milestones := daysUntil(t)

	body += fmt.Sprintf("\n\n🗓️ Days until:\n"+
		" 🌙 Next month: *%d days*\n"+
		" ☀️ Next year: *%d days*\n"+
		" 🔟 Next decade: *%d days*\n"+
		" 🕰️ Next century: *%d days*\n"+
		" 🛸 Next millennium: *%d days*\n",
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

	nasa, err := services.FetchNasaPhoto("today")
	if err == nil {
		body += fmt.Sprintf("\n🌌 NASA's Picture of the Day:\n[%s](%s)\n", nasa.Title, nasa.URL)
	} else {
		fmt.Println("Error fetching NASA photo:", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, body)
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

type TimeMilestones struct {
	DaysToNextMonth      int
	DaysToNextYear       int
	DaysToNextDecade     int
	DaysToNextCentury    int
	DaysToNextMillennium int
}

func daysUntil(from time.Time) TimeMilestones {
	// Normalize to start of day
	from = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())

	nextMonth := time.Date(from.Year(), from.Month()+1, 1, 0, 0, 0, 0, from.Location())
	nextYear := time.Date(from.Year()+1, 1, 1, 0, 0, 0, 0, from.Location())
	nextDecade := time.Date((from.Year()/10+1)*10, 1, 1, 0, 0, 0, 0, from.Location())
	nextCentury := time.Date((from.Year()/100+1)*100, 1, 1, 0, 0, 0, 0, from.Location())
	nextMillennium := time.Date((from.Year()/1000+1)*1000, 1, 1, 0, 0, 0, 0, from.Location())

	return TimeMilestones{
		DaysToNextMonth:      int(nextMonth.Sub(from).Hours() / 24),
		DaysToNextYear:       int(nextYear.Sub(from).Hours() / 24),
		DaysToNextDecade:     int(nextDecade.Sub(from).Hours() / 24),
		DaysToNextCentury:    int(nextCentury.Sub(from).Hours() / 24),
		DaysToNextMillennium: int(nextMillennium.Sub(from).Hours() / 24),
	}
}
