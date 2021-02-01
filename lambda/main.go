package main

import (
	"context"
	"fmt"
	"github.dxc.com/projects/aws-zerotohero/lambda/github"

	//"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) {
	github := github.NewGithub()
	err := github.Initialize("0f38dc7eba1c56cd8111677b26c534b86358d6e5")
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