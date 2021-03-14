package main

import (
	"context"
	"fmt"
	"github.com/idasilva/aws-zerotohero/lambda/k8s/client"
)

// GOOS=linux go build -o main
//zip main.zip main

const dsManifest = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80

`

func handler(ctx context.Context) error {

	//github := github.NewGithub()
	//err := github.Initialize(os.Getenv("GITHUB_ACCESS_TOKEN"))
	//if err != nil {
	//	return err
	//}
	//
	//err = github.NewVersion()
	//if err != nil {
	//	return err
	//}
	//
	//aws := codebuild.NewCodeBuild()
	//err = aws.Run()
	//if err != nil {
	//	return err
	//}
	//return nil
	client ,err := client.KubeConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = client.ApplyDeployment()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
func main() {
	handler(context.Background())
	//lambda.Start(handler)
}
