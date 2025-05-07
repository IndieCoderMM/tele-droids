package utils

import (
	"os"
	"strings"
	"time"
)

// ParseDate parses a date string in the format YYYY-MM-DD
func ParseDate(input string) (time.Time, error) {
	// DD-MM-YYYY format
	format := "2006-0102"
	// Parse the date string
	return time.Parse(format, strings.TrimSpace(input))
}

// Return env variable value as string
func GetEnvString(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return value
}
