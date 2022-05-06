package main

// kubectl get pod demo | go run .

import (
	"encoding/json"
	"hash"
	"hash/fnv"
	"os"

	"github.com/davecgh/go-spew/spew"

	v1 "k8s.io/api/core/v1"

	"k8s.io/klog/v2"
)

// DeepHashObject writes specified object to hash using the spew library
// which follows pointers and prints actual values of the nested objects
// ensuring the hash does not change when a pointer changes.
func DeepHashObject(hasher hash.Hash, objectToWrite interface{}) {
	hasher.Reset()
	printer := spew.ConfigState{
		Indent:         " ",
		SortKeys:       true,
		DisableMethods: true,
		SpewKeys:       true,
	}
	printer.Fprintf(hasher, "%#v", objectToWrite)
	// printer.Fprintf(os.Stderr, "got %s", objectToWrite)
}

// HashContainer returns the hash of the container. It is used to compare
// the running container with its desired spec.
// Note: remember to update hashValues in container_hash_test.go as well.
func HashContainer(container interface{}) uint64 {
	hash := fnv.New32a()
	// Omit nil or empty field when calculating hash value
	// Please see https://github.com/kubernetes/kubernetes/issues/53644
	containerJson, _ := json.Marshal(container)
	// klog.Infof("get container json: %s", containerJson)
	DeepHashObject(hash, containerJson)
	return uint64(hash.Sum32())
}

func main() {

	// can't use map[string]interface{} for v1.container

	pod := v1.Pod{}

	err := json.NewDecoder(os.Stdin).Decode(&pod)

	if err != nil {
		klog.Errorf("failed to decode stdin, please input json format: %v", err)
		os.Exit(1)
	}

	for _, cont := range pod.Spec.Containers {
		klog.Infof("get container %s, hash: %x", cont.Name, HashContainer(&cont))
	}

}
