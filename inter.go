package main

import (
	"fmt"
)

func inter(s1, s2 string) string {

	printed := make(map[rune]bool)
	result := ""

	for _, char := range s1 {
		if contains(s2, char) && !printed[char] {

			result += string(char)

			printed[char] = true
		}
	}

	return result + "\n"
}

func contains(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

func main() {
	fmt.Print(inter("padinton", "paqefwtdjetyiytjneytjoeyjnejeyj"))
	fmt.Print(inter("ddf6vewg64f", "twthgdwthdwfteewhrtag6h4ffdhsd"))
}
