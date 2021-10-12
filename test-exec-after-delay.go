package main

import (
	"os/exec"
	"time"
)

func main() {
	time.Sleep(time.Duration(5) * time.Second)
	c := exec.Command("/bin/ls")
	c.Start()
	c.Wait()
}
