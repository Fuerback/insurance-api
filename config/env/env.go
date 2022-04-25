package env

import "os"

func GetEnvWithDefaultAsString(envKey string, defaultVal string) string {
	val := os.Getenv(envKey)
	if val == "" {
		return defaultVal
	}
	return val
}
