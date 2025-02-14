package handlers

import (
	"backend/common"
	"fmt"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gofiber/fiber/v3"
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

	// Initialize the AWS SES client
	svc := common.NewAwsClient()

	// Create the SES input
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(req.Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data:    aws.String(req.Body),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(req.Subject),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(req.Sender),
	}

	// Send the email
	result, err := svc.SendEmail(input)
	if err != nil {
		// Log the error and return a 500 status code
		utils.WithStack(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("Failed to send email: %v", err),
		})
	}

	// Return the result
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Email sent successfully",
		"result":  result,
	})	
}

func TestHandler(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}