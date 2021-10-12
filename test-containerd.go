// https://github.com/containerd/containerd/blob/master/docs/getting-started.md
package main

import (
	"context"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {
	if err := redisExample(); err != nil {
		log.Fatal(err)
	}
}

//const s = "/run/containerd/containerd.sock"
// const s = "/proc/1964345/root/run/containerd/containerd.sock"
const s = "/proc/1964389/root/run/containerd/containerd.sock"

func redisExample() error {
	client, err := containerd.New(s)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "default")
	// image, err := client.Pull(ctx, "docker.io/library/redis:alpine", containerd.WithPullUnpack)
	image, err := client.Pull(ctx, "docker.io/library/bash:latest", containerd.WithPullUnpack)
	// image, err := client.Pull(ctx, "docker.io/library/bash:1.18.0", containerd.WithPullUnpack)
	if err != nil {
		return err
	}
	log.Printf("Successfully pulled %s image\n", image.Name())

	return nil
}
