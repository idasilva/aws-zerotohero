package client

import (
	s"github.com/idasilva/aws-zerotohero/lambda/aws/secret-manager"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


func makeConfig() (*rest.Config, error){

	manager := s.NewSecretManager()
	secret, err := manager.GetSecret()
	if err != nil{
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
	return cfg,nil
}
//KubeConfig
func KubeConfig() *kubernetes.Clientset{
	cfg , err := makeConfig()
	if err != nil {
		return nil
	}

	client ,err :=  kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil
	}
	return client
}