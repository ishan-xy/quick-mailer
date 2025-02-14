package common

import (
	"log"
	"os"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/joho/godotenv"
)
var IsDebug bool
func LoadEnv() {
	err := godotenv.Load()
	utils.WithStack(err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	IsDebug = os.Getenv("DEBUG") == "true"
}