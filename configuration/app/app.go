package app

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfiguration struct {
	AppHost string
	AppPort int
}

func GetenvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func NewConfiguration() *AppConfiguration {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	host := os.Getenv("API_HOST")
	port := GetenvInt("API_PORT", 8888)
	return &AppConfiguration{
		AppHost: host,
		AppPort: port,
	}
}
