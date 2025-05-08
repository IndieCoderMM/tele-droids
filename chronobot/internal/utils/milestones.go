package utils

import "time"

type TimeMilestones struct {
	DaysToNextMonth      int
	DaysToNextYear       int
	DaysToNextDecade     int
	DaysToNextCentury    int
	DaysToNextMillennium int
}

func DaysUntil(from time.Time) TimeMilestones {
	yy, mm, lo := from.Year(), from.Month(), from.Location()
	from = time.Date(yy, mm, from.Day(), 0, 0, 0, 0, lo)

	nextMonth := time.Date(yy, mm+1, 1, 0, 0, 0, 0, lo)
	nextYear := time.Date(yy+1, 1, 1, 0, 0, 0, 0, lo)
	nextDecade := time.Date((yy/10+1)*10, 1, 1, 0, 0, 0, 0, lo)
	nextCentury := time.Date((yy/100+1)*100, 1, 1, 0, 0, 0, 0, lo)

	daysUntillNextYear := int(nextYear.Sub(from).Hours() / 24)
	// Start from next year
	daysUntillMillennium := daysUntillNextYear
	currentYear := yy + 1
	nextMilYear := ((yy / 1000) + 1) * 1000
	for currentYear < nextMilYear {
		if IsLeapYear(currentYear) {
			daysUntillMillennium += 366
		} else {
			daysUntillMillennium += 365
		}
		currentYear++
	}

	return TimeMilestones{
		DaysToNextMonth:      int(nextMonth.Sub(from).Hours() / 24),
		DaysToNextYear:       daysUntillNextYear,
		DaysToNextDecade:     int(nextDecade.Sub(from).Hours() / 24),
		DaysToNextCentury:    int(nextCentury.Sub(from).Hours() / 24),
		DaysToNextMillennium: daysUntillMillennium,
	}
}
