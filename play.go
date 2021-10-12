package main

import (
	"fmt"
)

func get() (err error) {
	return fmt.Errorf("abc")
}

func main() {
	fmt.Print(get())
}
