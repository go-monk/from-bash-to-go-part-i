package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Millisecond * 100)
	fmt.Printf("\rfib(45) = %d\n", fib(45))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
