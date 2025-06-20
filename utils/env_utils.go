package utils

import (
	"os"
)

func GetEnvOrDefault(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvSplit(key string, fallback string) []string {
	env := GetEnvOrDefault(key, fallback)
	return SplitString(env)
}
