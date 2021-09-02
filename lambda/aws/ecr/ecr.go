package ecr

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/gopetbot/tidus/help"
	aws2 "github.com/idasilva/aws-zerotohero/lambda/aws"
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
		session.Must(session.NewSession(config.Configuration)), config.Configuration,
	)
	return &ContainerRegistry{
		ecr,
	}
}

//var imagesDetails []*ecr.ImageDetail
	//input := &ecr.ListImagesInput{
	//	RepositoryName: aws.String(""),
	//	RegistryId:     aws.String(""),
	//}
	//
	//pageIterator := 0
	//err := c.ListImagesPages(input, func(page *ecr.ListImagesOutput, lastPage bool) bool {
	//	pageIterator++
	//
	//	output, err := c.DescribeImages(&ecr.DescribeImagesInput{
	//		ImageIds:       page.ImageIds,
	//		RegistryId:     input.RegistryId,
	//		RepositoryName: input.RepositoryName,
	//	})
	//
	//	if err != nil {
	//		return true
	//	}
	//
	//	for _, i := range output.ImageDetails {
	//		if i.ImagePushedAt.After(*LatestPushedAt) {
	//			if len(i.ImageTags) > 0 {
	//				imagesDetails = append(imagesDetails,i)
	//				continue
	//			}
	//
	//			logrus.WithFields(logrus.Fields{
	//				"Digest": i.ImageDigest,
	//			}).Warn("image does not have a tag and will not be consider to validate")
	//		}
	//	}
	//
	//	return pageIterator <= 100
	//})
	//
	//if err != nil {
	//	return "", err
	//}
	//
	//sort.Slice(imagesDetails, func(i, j int) bool {
	//	return imagesDetails[i].ImagePushedAt.Before(*imagesDetails[j].ImagePushedAt)
	//})

	//return *imagesDetails[len(imagesDetails)-1].ImageTags[0], nil
