package main

import "fmt"

func main() {
	fmt.Printf("fib(45) = %d\n", fib(45))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
