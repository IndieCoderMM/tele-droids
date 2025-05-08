package utils_test

import (
	"chronobot/internal/utils"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	dateStrs := []string{
		"2023-10-01",
		"2023/10/01",
		"October 1, 2023",
		"Oct 1, 2023",
		"1 October 2023",
		"1 Oct 2023",
		"01-10-2023",
		"01/10/2023",
	}

	for _, dateStr := range dateStrs {
		t.Run(dateStr, func(t *testing.T) {
			parsed, err := utils.ParseDate(dateStr)
			if err != nil {
				t.Errorf("Failed to parse date: %v", err)
			}

			expected := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
			if parsed.Year() != expected.Year() || parsed.Month() != expected.Month() || parsed.Day() != expected.Day() {
				t.Errorf("Parsed date %v does not match expected date %v", parsed, expected)
			}
		})
	}

}
