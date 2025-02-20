package handlers

import (
	"backend/common"
	"backend/database"
	"log"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/gofiber/fiber/v3"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMail(c fiber.Ctx) error {
	var req database.EmailRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   utils.WithStack(err),
			"message": "Invalid request body",
		})
	}

	if req.Sender == "" || req.Recipient == "" || req.Subject == "" || req.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Sender, recipient, subject, and body are required",
		})
	}

	client, err := common.GetSMTPClient()
	if err != nil {
		return utils.WithStack(err)
	}
	email := mail.NewMSG()
	email.SetFrom(req.Sender).
		SetSubject(req.Subject).
		AddTo(req.Recipient).
		SetBody(mail.TextHTML, req.Body)
	email.SetDSN([]mail.DSN{mail.SUCCESS, mail.FAILURE}, false)

	if email.Error != nil {
		log.Printf("Failed to create email: %v", email.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   utils.WithStack(email.Error),
			"message": "Failed to create email message",
		})
	}

	err = email.Send(client)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   utils.WithStack(err),
			"message": "Failed to send email",
		})
	}

	email_struct := database.Email{
		Sender: req.Sender,
		Subject: req.Subject,
		Recipient: req.Recipient,
		TextBody: req.Body,
	}
	_, err = database.SentMailDB.InsertOne(c.Context(), email_struct)
	if err != nil {
		return utils.WithStack(err)
	}
	log.Println("Email sent successfully")
	return c.JSON(fiber.Map{
		"message": "Email sent successfully",
	})
}
