package ecr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/gopetbot/tidus/help"
	aws2 "github.dxc.com/projects/aws-zerotohero/lambda/aws"
	"os"
)

type ContainerRegistry struct {
	*ecr.ECR
}

type SortImageIds []*ecr.ImageIdentifier

func (c *ContainerRegistry) LatestImageTag() (string, error) {

	var imageIds []*ecr.ImageIdentifier

	params := &ecr.ListImagesInput{
		RepositoryName: aws.String(repoName),
		MaxResults:     aws.Int64(100),
		RegistryId:     aws.String(os.Getenv("AWS_ACCOUNT_ID")),
	}

	resp, err := c.ListImages(params)
	if err != nil {
		return help.Empty(), err
	}
	for _, imageID := range resp.ImageIds {

		imageIds = append(imageIds, imageID)
	}
	latestTag := *imageIds[len(imageIds)-1].ImageTag

	return latestTag, nil
}

//NewEcrInstance
func NewEcrInstance() *ContainerRegistry {
	config := aws2.NewRemote()

	ecr := ecr.New(
		session.Must(session.NewSession(config.Configuration)), aws.NewConfig().WithRegion("AWS_LAMBDA_REGION"),
	)
	return &ContainerRegistry{
		ecr,
	}
}
