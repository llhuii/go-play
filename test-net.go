package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://18.0.0.9:90"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get: %s", url, err)
	}
	defer resp.Body.Close()
}
