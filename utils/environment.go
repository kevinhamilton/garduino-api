package utils

import "os"

//GetEnv returns from OS environment variables, or from fallback.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
