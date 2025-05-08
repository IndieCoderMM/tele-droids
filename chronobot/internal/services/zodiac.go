package services

import "time"

// Returns the zodiac sign on the given date
func GetZodiac(t time.Time) string {
	month, day := t.Month(), t.Day()
	switch {
	case month == 3 && day >= 21 || month == 4 && day <= 19:
		return "Aries"
	case month == 4 && day >= 20 || month == 5 && day <= 20:
		return "Taurus"
	case month == 5 && day >= 21 || month == 6 && day <= 20:
		return "Gemini"
	case month == 6 && day >= 21 || month == 7 && day <= 22:
		return "Cancer"
	case month == 7 && day >= 23 || month == 8 && day <= 22:
		return "Leo"
	case month == 8 && day >= 23 || month == 9 && day <= 22:
		return "Virgo"
	case month == 9 && day >= 23 || month == 10 && day <= 22:
		return "Libra"
	case month == 10 && day >= 23 || month == 11 && day <= 21:
		return "Scorpio"
	case month == 11 && day >= 22 || month == 12 && day <= 21:
		return "Sagittarius"
	case month == 12 && day >= 22 || month == 1 && day <= 19:
		return "Capricorn"
	case month == 1 && day >= 20 || month == 2 && day <= 18:
		return "Aquarius"
	case month == 2 && day >= 19 || month == 3 && day <= 20:
		return "Pisces"
	}
	return "Unknown"
}

func GetChineseZodiac(year int) string {
	animals := []string{"Rat", "Ox", "Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Goat", "Monkey", "Rooster", "Dog", "Pig"}
	return animals[(year-4)%12]
}
