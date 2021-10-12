package main

import (
	"fmt"
)

type A struct{}

func (*A) func1() {
	fmt.Println("pointer func1")
}

func (A) func2() {
	fmt.Println("struct func2")
}

func main() {
	var a *A
	fmt.Println(a)
	(a.func1())
	(a.func2())

}
