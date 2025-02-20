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
	
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string

	API_Secret 	 string
}

func LoadConfig() (*Config, error) {

	return &Config{
		Port:        Getenv("PORT"),
		MongoURI:    Getenv("MONGO_URI"),
		DBName:      Getenv("MONGODB_DB"),
		JWTSecret:   Getenv("JWT_SECRET"),
		JWTExpiration: time.Hour * 24,
		CookieName:  "sessionID",

		SMTPHost:     Getenv("SMTP_HOST"),
		SMTPPort:     GetEnvAsInt("SMTP_PORT", 587),
		SMTPUsername: Getenv("SMTP_USERNAME"),
		SMTPPassword: Getenv("SMTP_PASSWORD"),
		API_Secret:   Getenv("API_Secret"),
	}, nil
}
