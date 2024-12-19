package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key, defaultVal string) string {
	if err := godotenv.Load(); err != nil {
		return defaultVal
	}
	return os.Getenv(key)
}

func GetInt(key string, defaultVal int) int {
	if err := godotenv.Load(); err != nil {
		return defaultVal
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return valAsInt
}
