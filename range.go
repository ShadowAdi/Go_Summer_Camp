package main

import "fmt"

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonaaci() {
	n := 10
	a, b := 0, 1
	fmt.Println("Fibonacci Series:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		next := a + b
		a = b
		b = next
	}
	fmt.Println()
}

func main() {
	fibonaaci()
}
