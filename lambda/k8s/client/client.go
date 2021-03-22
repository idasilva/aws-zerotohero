package client

import (
	"context"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/gopetbot/tidus/help"
	ecr2 "github.com/idasilva/aws-zerotohero/lambda/aws/ecr"
	secret_manager2 "github.com/idasilva/aws-zerotohero/lambda/aws/secret-manager"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

type K8s struct {
	kubectl *kubernetes.Clientset
	logger  *help.Logging
	box *packr.Box
}
//ApplyDeployment
func (k *K8s) ApplyDeployment() error {

	k.logger.Info("its will try put deployment on cluster..")

	tag, err:= ecr2.NewEcrInstance().LatestImageTag()
	if err != nil{
		k.logger.WithFields(logrus.Fields{
			"tag":tag,
		}).Info("it was not possible get latest image from ecr...")
		return err
	}

	deployment, err := k.ReadDeploymentFromYaml()
	if err != nil{
		k.logger.Info("it was not possible read yaml....")
		return err
	}

	image := fmt.Sprintf(os.Getenv("AWS_ECR_URI")+tag)

    deployment.Spec.Template.Spec.Containers[0].Image =  image
	clt := k.kubectl.AppsV1().Deployments(k.NameSpace())
	if err != nil {
		return err
	}

	deployment, err = clt.Create(context.TODO(), deployment, v1.CreateOptions{})
	if err != nil {
		return err
	}

	k.logger.WithFields(logrus.Fields{
		"image": image,
		"status":      deployment.Status,
		"name":        deployment.Name,
		"api-version": deployment.APIVersion,
		"namespace":   deployment.Namespace,
	}).Info("deployment done with success...")
	return nil
}
//ReadDeploymentFromYaml
func(k *K8s) ReadDeploymentFromYaml() (*appsv1.Deployment, error){
	yamlFile, err := k.box.FindString("deployment.yaml")
	if err != nil{
		return nil, err
	}

    deployment := &appsv1.Deployment{}
	dec := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	_, _, err = dec.Decode([]byte(yamlFile), nil, deployment)
	if err != nil{
		return nil, err
	}
	return deployment, nil
}
//NameSpace
func (k *K8s) NameSpace() string {
	return "default"
}

func makeConfig() (*rest.Config, error) {
	manager := secret_manager2.NewSecretManager()
	secret, err := manager.GetSecret()
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.NewClientConfigFromBytes([]byte(secret))
	if err != nil {
		return nil, err
	}

	cfg, err := config.ClientConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

//KubeConfigF
func KubeConfigF() (*K8s, error) {
	cfg, err := makeConfig()
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &K8s{
		kubectl: client,
		logger:  help.NewLog(),
		box : packr.New("someBoxName", "./assets"),
	}, nil
}
