package main

import (
	"fmt"
)

func main() {
	var nilD []int
	for _ = range nilD {

	}
	fmt.Println("len(nil): %d", len(nilD))
}
