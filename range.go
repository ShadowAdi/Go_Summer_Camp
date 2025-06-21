package main

import "fmt"

func main() {
	ans := "aditya"
	mpp := make(map[rune]int)
	for _, ch := range ans {
		mpp[ch]++
	}

	for k, v := range mpp {
		fmt.Printf("%q: %d\n", k, v)
	}
}
