package main

import (
	"chronobot/utils"
	"fmt"
	"time"
)

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
