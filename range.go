package main

import "fmt"

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Before:", arr)
	reverse(arr)
	fmt.Println("After:", arr)
}
