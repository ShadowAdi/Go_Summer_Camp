package main

import "fmt"

func add_func(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func add(x, y int) int {
	return x + y
}
func main() {
	a := add_func(3, 4, add)
	fmt.Println(a)

}
