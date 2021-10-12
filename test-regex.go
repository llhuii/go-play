package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := `\{2,5}?`
	p, err := regexp.Compile(r)
	_ = p
	fmt.Printf("==>: %v\n", err)
	m := p.MatchString("9999")
	m = p.MatchString("{2,5}")
	fmt.Println(m, err)

}
