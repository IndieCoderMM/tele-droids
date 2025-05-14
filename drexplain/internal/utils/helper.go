package utils

import (
	"os"
)

// Return env variable value as string
func GetEnvString(key, fallback string) (string, bool) {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback, false
	}

	return value, true
}
