package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"os"
)

type AwsConfig struct {
	Configuration *aws.Config
}

func NewRemote() *AwsConfig {

	aws := &aws.Config{
		Region: aws.String(os.Getenv("AWS_LAMBDA_REGION")),
	}
	return &AwsConfig{
		aws,
	}
}
