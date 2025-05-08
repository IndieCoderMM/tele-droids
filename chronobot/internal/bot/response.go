package bot

import (
	"chronobot/internal/services"
	"fmt"
	"time"
)

func BuildResponse(t time.Time) string {
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

	event, birth, err := services.FetchOnThisDay(t.Month(), t.Day())
	if err != nil {
		fmt.Println("Error fetching events:", err)
	}
	nasa, err := services.FetchNasaPhoto(t.Format("2006-01-02"))
	if err != nil {
		fmt.Println("Error fetching NASA photo:", err)
	}

	return fmt.Sprintf("🕰️ *%s* — that's %d days ago!\n\n📆 It was a *%s*\n♈ People born on this day are *%s*\n🐲 In Chinese zodiac, they'd be a *%s*\n\n📜 *Back in time, on this day:*\n%s\n\n🎂 *Famous birthdays:*\n%s\n\n📸 *NASA photo of the day:*\n[%s](%s)",
		t.Format("2006-01-02"), daysAgo, weekday, zodiac, chineseZodiac, event, birth, nasa.Title, nasa.URL)
}
