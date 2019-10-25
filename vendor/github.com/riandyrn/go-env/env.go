package env

import (
	"os"
	"strconv"
)

// GetString is used for fetching string value from env variables.
// Returns empty string if key not exist.
func GetString(key string) string {
	return os.Getenv(key)
}

// GetInt is used for fetching int value from env variables.
// Returns `0` if key not exist or has invalid value.
func GetInt(key string) int {
	v, _ := strconv.Atoi(GetString(key))
	return v
}

// GetBool is used for fetching bool value from env varibles.
// Returns `false` if key is not exist or has invalid value.
func GetBool(key string) bool {
	v, _ := strconv.ParseBool(GetString(key))
	return v
}
