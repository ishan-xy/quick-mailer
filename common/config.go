package common

import (
	"time"
)

type Config struct {
	Port        string
	MongoURI    string
	DBName      string
	JWTSecret   string
	JWTExpiration time.Duration
	CookieName 	string
	
}

func LoadConfig() (*Config, error) {

	return &Config{
		Port:        Getenv("PORT"),
		MongoURI:    Getenv("MONGO_URI"),
		DBName:      Getenv("MONGODB_DB"),
		JWTSecret:   Getenv("JWT_SECRET"),
		JWTExpiration: time.Hour * 24,
		CookieName:  "sessionID",
	}, nil
}
