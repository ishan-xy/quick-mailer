package handlers

import (
	"backend/common"
	"log"

	"github.com/gofiber/fiber/v3"
	mail "github.com/xhit/go-simple-mail/v2"
)

type EmailRequest struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func SendMail(c fiber.Ctx) error {
	var req EmailRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
		})
	}

	if req.Sender == "" || req.Recipient == "" || req.Subject == "" || req.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Sender, recipient, subject, and body are required",
		})
	}

	client := common.GetSMTPClient()
	email := mail.NewMSG()
	email.SetFrom(req.Sender).
		SetSubject(req.Subject).
		AddTo(req.Recipient).
		SetBody(mail.TextHTML, req.Body)
	email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, false)

	if email.Error != nil {
		log.Printf("Failed to create email: %v", email.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create email message",
		})
	}

	err := email.Send(client)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to send email",
		})
	}

	log.Println("Email sent successfully")
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Email sent successfully",
	})
}

func TestHandler(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}