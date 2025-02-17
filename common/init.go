package common

import (
	utils "github.com/ItsMeSamey/go_utils"
	"log"
)

var Cfg *Config

func init() {
	loadEnv()
	var err error
	Cfg, err = LoadConfig()
	if err != nil {
		log.Fatal(utils.WithStack(err))
	}
	log.Println("Configuration loaded successfully:", Cfg)
}