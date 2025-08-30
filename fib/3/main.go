package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fib()
	for i := 0; i <= 45; i++ {
		fmt.Printf("fib(%d) = %d\n", i, <-c)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func fib() <-chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}
