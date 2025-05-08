package utils

import (
	"os"
	"strings"
	"time"
)

// ParseDate parses a date string in the format YYYY-MM-DD
func ParseDate(input string) (time.Time, error) {
	// DD-MM-YYYY format
	format := "2006-01-02"
	// Parse the date string
	return time.Parse(format, strings.TrimSpace(input))
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
