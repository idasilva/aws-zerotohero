package main

import (
	"context"
	"fmt"
	"github.com/gopetbot/tidus/help"
	"github.com/idasilva/aws-zerotohero/lambda/aws/codebuild"
	"github.com/idasilva/aws-zerotohero/lambda/aws/sns"
	"github.com/idasilva/aws-zerotohero/lambda/github"
	client2 "github.com/idasilva/aws-zerotohero/lambda/k8s/client"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

// GOOS=linux go build -o main
//zip main.zip main

func handler(ctx context.Context) error {
	logger := help.NewLog()

	github := github.NewGithub()
	err := github.Initialize(os.Getenv("GITHUB_ACCESS_TOKEN"))
	if err != nil {
		return err
	}

	err = github.NewVersion()
	if err != nil {
		return err
	}

	logger.Info("update version")

	aws := codebuild.NewCodeBuild()
	err = aws.Run()
	if err != nil {
		return err
	}

	client,err:= client2.KubeConfigF()
	if err != nil {
		return err
	}

	err = client.ApplyDeployment()
	if err != nil {
		return err
	}
	logger.Info("deploy...")

	sns := sns.NewSNS()
	err = sns.PublishMessage(fmt.Sprint("deploy feito com sucesso!"),os.Getenv("AWS_SNS_TOPIC"))
	if err != nil{
		return err
	}

	logger.Info("finish lambda..")
	return nil
}
func main() {
	//handler(context.Background())
	lambda.Start(handler)
}
