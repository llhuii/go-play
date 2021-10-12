package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	clientset, err := kubernetes.NewForConfig(config)
	_ = err
	node, err = clientset.CoreV1().Nodes().List(context.TODO(),
		metav1.ListOptions{
			// TimeoutSeconds: &t,
		})
	fmt.Printf("xxx %+v %+v", node, err)
}
