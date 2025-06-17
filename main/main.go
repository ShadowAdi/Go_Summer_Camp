package main

import "fmt"

func main() {
	s := [6]int{1, 2, 3, 4, 5, 7}
	fmt.Println(s)
	a := make([]int, 2)
	a = append(a, 9)
	fmt.Println(a)

}
