package codebuild

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/gopetbot/tidus/help"
	_ "github.com/gopetbot/tidus/help"
	"github.com/sirupsen/logrus"
	config "github.dxc.com/projects/aws-zerotohero/lambda/aws"
	"os"
)

var (
	projectName = os.Getenv("PROJECT_NAME")
)

type Project struct {
	name   string
	logger *help.Logging
	client *codebuild.CodeBuild
}

func (p *Project) Run() error {
	_ = p.inputProjectName()
	_, err := p.client.StartBuild(&codebuild.StartBuildInput{ProjectName: aws.String(p.name)})
	if err != nil {
		return err
	}
	p.logger.WithFields(logrus.Fields{
		"projectName": p.name,
	}).Info("started build...")

	return nil
}
func (p *Project) inputProjectName() error {
	if help.IsEmpty(projectName) {
		p.logger.Info("project name required!...")
		return errors.New("was not possible build a project without a name")
	}
	p.name = projectName
	return nil
}

func NewCodeBuild() *Project {

	remote := config.NewRemote()

	var client = codebuild.New(
		session.Must(
			session.NewSession(remote.Configuration)),
		aws.NewConfig().WithRegion(os.Getenv("AWS_LAMBDA_REGION")),
	)

	return &Project{
		client: client,
		logger: help.NewLog(),
	}
}
