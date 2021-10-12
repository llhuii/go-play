package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var p *int
	b, e := json.Marshal(p)
	if e != nil {
		fmt.Printf("Marshal error %v\n", e)
		return
	}
	fmt.Printf("Marshal string: %s\n", b)
}
