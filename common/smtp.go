package common

import (
	"crypto/tls"
	"log"
	"strconv"
	"time"

	utils "github.com/ItsMeSamey/go_utils"
	mail "github.com/xhit/go-simple-mail/v2"
)

func GetSMTPClient() *mail.SMTPClient {
	server := mail.NewSMTPClient()

	server.Host = Getenv("SMTP_HOST")
	port, err := strconv.Atoi(Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}
	server.Port = port
	server.Username = Getenv("SMTP_USERNAME")
	server.Password = Getenv("SMTP_PASSWORD")
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	client, err := server.Connect()
	if err != nil {
		utils.WithStack(err)
		log.Fatalf("Failed to connect to SMTP server: %v", err)
	}

	return client
}
