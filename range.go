package main

import "fmt"

func addT(a, b int64) int64 {
	return a + b
}

func main() {
	a := 12
	if a > 10 {
		fmt.Printf("The Value of a is: %d ", a)
	} else {
		fmt.Printf("The Value of a is less than 10: %d ", a)
	}

}
