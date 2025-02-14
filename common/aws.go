package common

import (
	"os"

	utils "github.com/ItsMeSamey/go_utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func NewAwsClient() *ses.SES {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		),
	})
	utils.WithStack(err)
	return ses.New(sess)
}