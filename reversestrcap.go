package main

import (
	"fmt"
	"strings"
)

// Function to process each word in the input string
func reverseStrCap(input string) string {
	words := strings.Fields(input) // Split the string into words
	var result []string

	// Process each word
	for _, word := range words {
		if len(word) > 0 {
			// Convert the whole word to lowercase
			word = strings.ToLower(word)
			// Capitalize the last letter
			word = word[:len(word)-1] + strings.ToUpper(string(word[len(word)-1]))
		}
		result = append(result, word)
	}

	// Join the words back into a single string and return it
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(reverseStrCap("First SMALL TesT"))
}
