package main

import (
	"fmt"
	"math/rand"
	"time"
)

func g(m time.Duration) time.Duration {
	r := 1 + rand.Float64()
	return time.Duration(r * float64(m.Nanoseconds()))
}
func main() {
	var nilMap map[string]string
	rand.Seed(time.Now().UnixNano())
	m := 30 * time.Second
	fmt.Println(" ===", g(m), nilMap["aaa"])
	fmt.Println(" ===", g(m))
}
