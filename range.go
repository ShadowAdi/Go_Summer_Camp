package main

import "fmt"

func addT(a, b int64) int64 {
	return a + b
}

func main() {
	result := addT(12, 10)
	fmt.Println("Sum:", result)

}
