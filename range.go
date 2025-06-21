package main

import "fmt"

func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	arr := make([]int, 5)
	for i, _ := range arr {
		arr[i] = i + 1
	}

	arr = append(arr, 4)

	moreNumbers := []int{8, 9, 10}
	arr = append(arr, moreNumbers...)
	fmt.Println("After appending another slice:", arr)

}
