package main

import (
	"context"
	"io"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

func ReadAll(r io.Reader) ([]byte, error) {
	b := make([]byte, 0, 512)
	// b = make([]byte, 0, 8388608)
	for {
		if len(b) == cap(b) {
			// Add more capacity (let append pick how much).
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return b, err
		}
	}
}

func main() {
	f := os.Getenv("KUBECONFIG")
	if len(f) == 0 {
		f = os.Getenv("HOME") + "/.kube/config"
	}
	config, err := clientcmd.BuildConfigFromFlags("", f)
	clientset, err := kubernetes.NewForConfig(config)
	_ = err
	d := clientset.DiscoveryClient
	request := d.RESTClient().Get().AbsPath("/openapi/v2")
	{
		klog.Infof("before get openapi Stream to File()")
		stream, err := request.Stream(context.TODO())
		if err != nil {
			klog.Warningf("failed to get openapi Stream(): %v", err)

		} else {
			file, _ := os.CreateTemp("/tmp", "dis-xxx")
			io.Copy(file, stream)
			f.Close()
			defer os.Remove(f.Name()) // clean up

			klog.Infof("after get openapi Stream to File ()")
		}
	}
	{
		klog.Infof("before get openapi Stream to Memory()")
		stream, err := request.Stream(context.TODO())
		if err != nil {
			klog.Warningf("failed to get openapi Stream(): %v", err)

		} else {
			bytes, _ := ReadAll(stream)
			klog.Infof("after get openapi Stream to Memory (): %d", len(bytes))
		}
	}
	{
		klog.Infof("before get openapi DoRaw()")
		data, _ := request.DoRaw(context.TODO())
		klog.Infof("after get openapi DoRaw(): %d", len(data))
	}
	{
		data, _ := request.Do(context.TODO()).Raw()
		klog.Infof("after get openapi Do().Raw(): %d", len(data))
	}
	{
		klog.Infof("before get openapi OpenAPISchema()")
		doc, err := d.OpenAPISchema()
		if err != nil {
			klog.Infof("failed to get openapi:%+v", err)
		} else {
			klog.Infof("after get openapi OpenAPISchema: %s", doc.Swagger)

		}
	}
}
