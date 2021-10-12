package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		defer func() {
			fmt.Println("i", recover())
			fmt.Println("j", recover())
		}()
		panic(1)
	}()
	panic(2)
}
