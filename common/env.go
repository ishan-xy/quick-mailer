package common

import (
	"log"
	"os"
	"strconv"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/joho/godotenv"
)

var IsDebug bool

func loadEnv() {
	err := godotenv.Load()
	utils.WithStack(err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	IsDebug = os.Getenv("DEBUG") == "true"
}

func Getenv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Missing environment variable: %s", key)
	}
	return value
}

func GetEnvAsInt(key string, defaultValue int) (val int, err error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue, nil
	}
	return strconv.Atoi(value)
}

func GetEnvAsBool(key string, defaultValue bool) (val bool) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value == "true"
}