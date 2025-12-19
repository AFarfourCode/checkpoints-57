package main

import "fmt"

func FifthAndSkip(str string) string {
	// Case 1: empty string
	if str == "" {
		return "\n"
	}

	// Count non-space characters
	count := 0
	for _, ch := range str {
		if ch != ' ' {
			count++
		}
	}

	// Case 2: less than 5 valid characters
	if count < 5 {
		return "Invalid Input\n"
	}

	result := ""
	charCount := 0  // counts valid characters (non-space)
	groupCount := 0 // counts characters inside a group of 5

	for _, ch := range str {
		if ch == ' ' {
			continue
		}

		charCount++

		// Skip every 6th character
		if charCount%6 == 0 {
			continue
		}

		result += string(ch)
		groupCount++

		// Add space after every 5 characters (if more characters remain)
		if groupCount == 5 {
			result += " "
			groupCount = 0
		}
	}

	// Remove trailing space if exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
