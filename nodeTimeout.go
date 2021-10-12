package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	host := "182.0.0.1"
	t := int64(5)
	_ = t
	if len(os.Args) > 1 {

		host = os.Args[1]
	}
	if len(os.Args) > 2 {
		i, _ := strconv.Atoi((os.Args[2]))
		t = int64(i)
	}
	config, err := clientcmd.BuildConfigFromFlags("http://"+host+":443", "/root/.kube/config")
	clientset, err := kubernetes.NewForConfig(config)
	_ = err
	_, err = clientset.CoreV1().Nodes().List(context.TODO(),
		metav1.ListOptions{
			// TimeoutSeconds: &t,
		})
	fmt.Printf("xxx %+v", err)
}
