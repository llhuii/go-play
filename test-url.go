package main

import (
	"fmt"
	"net/url"
)

var ss = [...]string{
	// s3 compatibile storage
	"s3",

	// http server
	"http", "https",

	// hostpath of node, for compatibility only
	// "/opt/data/model.pb"
	"",

	// the local path of worker-container
	"file",
}

var us = [...]string{
	// s3 compatibile storage
	"s3",

	// http server
	"http", "https",

	// hostpath of node, for compatibility only
	// "/opt/data/model.pb"
	"",

	// the local path of worker-container
	"file",
}

func main() {
	for _, uri := range []string{
		"https://aaa.b/c/d/e.tar.gz?p=9",
		"/aaa.b/c/d/e.tar.gz?p=9",
		"s3://aaa.b/c/d/e.tar.gz?p=9",
		"file://aaa.b/c/d/e%30.tar.gz?p=9",
	} {
		u, err := url.Parse(uri)
		fmt.Println("parsing", uri)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("\t scheme", u.Scheme)
		fmt.Println("\t escaped path", u.EscapedPath())
		fmt.Println("\t host", u.Host)
		fmt.Println("\t path", u.Path)
		u.Path = "aaaa"
	}

}
