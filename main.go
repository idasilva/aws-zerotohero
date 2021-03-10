package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


func main()  {

	config, err := clientcmd.NewClientConfigFromBytes([]byte(""))
	if err != nil {
		return
	}
	newConfig, err := config.ClientConfig()
	if err != nil {
		return
	}
	client ,_ :=  kubernetes.NewForConfig(newConfig)
	fmt.Println(client)
}
