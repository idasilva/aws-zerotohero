package main

import (
	"context"
	"fmt"
	"github.dxc.com/projects/aws-zerotohero/lambda/github"
	"os"

	//"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) {
	github := github.NewGithub()
	err := github.Initialize(os.Getenv("GITHUB_ACCESS_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	err = github.UpdateVersion()
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	handler(context.Background())
	//lambda.Start(handler)
}