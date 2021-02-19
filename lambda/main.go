package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.dxc.com/projects/aws-zerotohero/lambda/aws/codebuild"
	"github.dxc.com/projects/aws-zerotohero/lambda/github"
	"os"
)

// GOOS=linux go build -o main
//zip main.zip main

func handler(ctx context.Context) error {

	github := github.NewGithub()
	err := github.Initialize(os.Getenv("GITHUB_ACCESS_TOKEN"))
	if err != nil {
		return err
	}

	err = github.NewVersion()
	if err != nil {
		return err
	}

	aws := codebuild.NewCodeBuild()
	aws.Run()
	if err != nil {
		return err
	}
	return nil

}
func main() {
	lambda.Start(handler)
}
