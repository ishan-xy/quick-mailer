package main

import (
	"fmt"
	"os"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/joho/godotenv"
)


func awsmail() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	recipients := os.Getenv("RECIPIENTS")
	sender := os.Getenv("SENDER")
	
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(
			"YOUR_ACCESS_KEY_ID",
			"YOUR_SECRET_ACCESS_KEY",
			"",
		),
    })
    if err != nil {
        fmt.Println("Failed to create session:", err)
        return
    }

    svc := ses.New(sess)

    input := &ses.SendEmailInput{
        Destination: &ses.Destination{
            ToAddresses: []*string{
                aws.String(recipients),
            },
        },
        Message: &ses.Message{
            Body: &ses.Body{
                Text: &ses.Content{
                    Data:    aws.String("This is the email body."),
                    Charset: aws.String("UTF-8"),
                },
            },
            Subject: &ses.Content{
                Data:    aws.String("Test Email via Amazon SES"),
                Charset: aws.String("UTF-8"),
            },
        },
        Source: aws.String(sender),
    }

    result, err := svc.SendEmail(input)
    if err != nil {
        fmt.Println("Error sending email:", err)
        return
    }

    fmt.Printf("Email sent! Message ID: %s\n", *result.MessageId)
}