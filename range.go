package main

import "fmt"

func main() {
	ans := "aditya"
	ans1 := "dityaa"

	if isAnagram(ans, ans1) {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal")
	}
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, ch := range a {
		m1[ch]++
	}
	for _, ch := range b {
		m2[ch]++
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
