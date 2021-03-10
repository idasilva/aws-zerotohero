package secret_manager

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	secret "github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/gopetbot/tidus/help"
	aws2 "github.com/idasilva/aws-zerotohero/lambda/aws"
	"github.com/sirupsen/logrus"
)

type Secret struct {
	Manager *secret.SecretsManager

	logger       *help.Logging
	SecretID     *string
	VersionStage *string
}

func (s *Secret) GetSecret()( string,error) {

	result, err := s.Manager.GetSecretValue(&secret.GetSecretValueInput{
		SecretId:     s.SecretID,
		VersionStage: s.VersionStage,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			Is := Manager(aerr.Code())
			if help.IsEmpty(string(Is)) {
				s.logger.WithFields(logrus.Fields{
					"err": Is,
				}).Info("error not exist in map...")
				return help.Empty(), err
			}
			s.logger.Infof("error: %s", err)
			return  help.Empty(),err

		} else {
			s.logger.Info("cast err to awserr.error to get the code and...")
			return  help.Empty(),err

		}
		return  help.Empty(),err

	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, decodedBinarySecret string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return  help.Empty(),err

		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
	}
	fmt.Println(secretString, decodedBinarySecret)
	// Your code goes here.

	return  secretString, nil
}

//NewSecretManager create a secrets manager client
func NewSecretManager() Secret {
	config := aws2.NewRemote()
	manager := secret.New(session.Must(session.NewSession(config.Configuration)),
		config.Configuration)

	return Secret{
		Manager:      manager,
		SecretID:     aws.String("AWS_SECRET_NAME"),
		VersionStage: aws.String("AWSCURRENT"),
	}
}
