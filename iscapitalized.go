package main

import "fmt"

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	runes := []rune(s)
	if runes[0] >= 'a' && runes[0] <= 'z' {
		return false
	}
	for i := 1; i < len(runes); i++ {
		if runes[i-1] == ' ' {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
