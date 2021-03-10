package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


func secretManager() (*rest.Config, error){
	config, err := clientcmd.NewClientConfigFromBytes([]byte(""))
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
	cfg , err := secretManager()
	if err != nil {
		return nil
	}

	client ,err :=  kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil
	}
	return client
}