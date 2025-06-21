package main

import (
	"fmt"
	"strconv"
)

func main() {
	words := []string{"cat", "go", "banana", "a", "hello", "hat", "to", "do"}
	grouped := make(map[string][]string)
	for _, word := range words {
		lengthKey := strconv.Itoa(len(word))
		grouped[lengthKey] = append(grouped[lengthKey], word)
	}

	for length, wordList := range grouped {
		fmt.Printf("Length %s: %v\n", length, wordList)
	}
}
