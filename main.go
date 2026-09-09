package main

import (
	"fmt"
	"strings"

)

func countWords(text string) map[string]int {
	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
} 


func main() {
	text := "the wolf protects the child and the child remembers the wolf"

	counts := countWords(text)

	fmt.Println(counts)

}