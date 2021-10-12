package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a, _ := filepath.Split("/")
	fmt.Println(a)

}
