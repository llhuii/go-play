package main

import (
	"fmt"
)

func main() {
	defer func() {
		defer func() {
			fmt.Println(recover())
		}()
		panic(1)
	}()
	defer func() {
		fmt.Println(recover())
	}()
	panic(2)
}
