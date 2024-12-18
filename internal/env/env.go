package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetString(key, defaultVal string) string {
	if err := godotenv.Load(); err != nil {
		return defaultVal
	}
	return os.Getenv(key)
}
