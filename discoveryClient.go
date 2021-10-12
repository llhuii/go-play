package main

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	clientset, err := kubernetes.NewForConfig(config)
	d := clientset.DiscoveryClient
	{
		klog.Infof("before get openapi after DoRaw()")
		o := d.RESTClient().Get().AbsPath("/openapi/v2")
		klog.Infof("before get openapi after DoRaw()")
		data, _ := o.DoRaw(context.TODO())
		klog.Infof("get openapi after DoRaw(): %d", len(data))
	}
	{
		data, _ := d.RESTClient().Get().AbsPath("/openapi/v2").Do(context.TODO()).Raw()
		klog.Infof("get openapi after Do().Raw(): %d", len(data))
	}
	doc, err := d.OpenAPISchema()
	if err != nil {
		klog.Infof("failed to get openapi:%+v", err)
	} else {
		klog.Infof("get openapi: %s", doc.Swagger)

	}
}
