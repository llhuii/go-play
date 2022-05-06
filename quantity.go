package main

import (
        "fmt"

        "k8s.io/apimachinery/pkg/api/resource"
)

func main() {
        limit, _ := resource.ParseQuantity("110m")
        req, _ := resource.ParseQuantity("0")
        fmt.Printf("req is zero %s\n", req.IsZero())
        req.SetMilli(limit.ScaledValue(2 - 3))
        fmt.Printf("limit %d\n", limit.Value())
        fmt.Printf("got scale %d\n", limit.ScaledValue(2))
        fmt.Printf("got req %s", req.String())

}
