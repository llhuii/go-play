package main

import (
	"fmt"
	"sync"
)

type C struct {
	c chan int
}

func test1() {
	var c C
	c.c = make(chan int)
	go func(c C) {
		o := c.c
		for i := 0; i < 9; i++ {
			o <- i
		}
		close(o)

	}(c)
	o := c.c

	for i := range o {
		fmt.Println(i)
	}

}

func closeAndRange() {
	size := 5
	channels := make(chan int, size)
	wg := sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			channels <- i*2 + 1
		}(i)
	}
	wg.Wait()
	close(channels)
	for i := range channels {
		fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	closeAndRange()
}
