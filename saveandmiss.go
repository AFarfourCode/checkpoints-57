package main

import "fmt"

// SaveAndMiss processes the string based on the provided chunk size `num`.
// It saves the first chunk, skips the second, saves the third, and so on.
func SaveAndMiss(arg string, num int) string {
	// If num is less than or equal to 0, return the original string
	if num <= 0 {
		return arg
	}

	var result string

	// Loop through the string in chunks of `num`
	for i := 0; i < len(arg); i += num * 2 {
		// Save the first `num` characters
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}
		result += arg[i:end]
	}

	// Return the string with saved characters
	return result
}

func main() {
	// Test cases
	fmt.Println(SaveAndMiss("123456789", 3))                  // 123789
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))  // abcghimnostuz
	fmt.Println(SaveAndMiss("", 3))                           // ""
	fmt.Println(SaveAndMiss("hello you all ! ", 0))           // hello you all !
	fmt.Println(SaveAndMiss("what is your name?", 0))         // what is your name?
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5)) // go Exercise Save and Miss
}
