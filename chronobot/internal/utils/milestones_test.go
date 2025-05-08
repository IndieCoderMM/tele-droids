package utils_test

import (
	"chronobot/internal/utils"
	"testing"
	"time"
)

func TestDaysUntil(t *testing.T) {
	now := time.Date(2025, 5, 8, 0, 0, 0, 0, time.UTC)

	result := utils.DaysUntil(now)

	expectedNextMonth := 24
	expectedNextYear := 238
	expectedNextDecade := 1699
	expectedNextCentury := 27266
	expectedMillennium := 355984

	if result.DaysToNextMonth != expectedNextMonth {
		t.Errorf("DaysToNextMonth = %d, want %d", result.DaysToNextMonth, expectedNextMonth)
	}
	if result.DaysToNextYear != expectedNextYear {
		t.Errorf("DaysToNextYear = %d, want %d", result.DaysToNextYear, expectedNextYear)
	}
	if result.DaysToNextDecade != expectedNextDecade {
		t.Errorf("DaysToNextDecade = %d, want %d", result.DaysToNextDecade, expectedNextDecade)
	}
	if result.DaysToNextCentury != expectedNextCentury {
		t.Errorf("DaysToNextCentury = %d, want %d", result.DaysToNextCentury, expectedNextCentury)
	}
	if result.DaysToNextMillennium != expectedMillennium {
		t.Errorf("DaysToNextMillennium = %d, want %d", result.DaysToNextMillennium, expectedMillennium)
	}
}
