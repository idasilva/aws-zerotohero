package main

import (
	"context"
	"fmt"
	"github.dxc.com/projects/aws-zerotohero/lambda/github"

	//"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) {
	github := github.NewGithub()
	err := github.Initialize("accessKey")
	if err != nil {
		fmt.Println(err)
	}

	err = github.Version()
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	handler(context.Background())
	//lambda.Start(handler)
}