package utils

import (
	"fmt"
	"os"
	"time"
)

// ParseDate parses a date string in different formats
func ParseDate(input string) (time.Time, error) {
	// Define the date formats to try
	dateFormats := []string{
		"2006-01-02",      // YYYY-MM-DD
		"2006/01/02",      // YYYY/MM/DD
		"02-01-2006",      // DD-MM-YYYY
		"02/01/2006",      // DD/MM/YYYY
		"January 2, 2006", // Month DD, YYYY
		"Jan 2, 2006",     // Month DD, YYYY
		"2 January 2006",  // DD Month YYYY
		"2 Jan 2006",      // DD Month YYYY
	}

	for _, format := range dateFormats {
		t, err := time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", input)
}

func IsLeapYear(year int) bool {
	// Check if the year is divisible by 4 and not divisible by 100, or divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

// Return env variable value as string
func GetEnvString(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}
