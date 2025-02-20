package common

import (
	"crypto/tls"
	"log"
	"time"

	utils "github.com/ItsMeSamey/go_utils"
	mail "github.com/xhit/go-simple-mail/v2"
)

func GetSMTPClient() (*mail.SMTPClient, error) {
	server := mail.NewSMTPClient()

	server.Host = Cfg.SMTPHost
	server.Port = Cfg.SMTPPort
	server.Username = Cfg.SMTPUsername
	server.Password = Cfg.SMTPPassword
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	client, err := server.Connect()
	if err != nil {
		log.Printf("Failed to connect to SMTP server: %v", utils.WithStack(err))
	}

	return client ,err
}
