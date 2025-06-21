package main

import "fmt"

func main() {
	var arr1 = [3]int{10, 11, 1}
	var maxNum int
	for _, v := range arr1 {
		maxNum = max(maxNum, v)
	}
	fmt.Println("The value is %d: ", maxNum)

}
