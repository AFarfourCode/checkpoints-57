package main

import (
	"fmt"
)

func lastword(s string) string {
	if s == "" {
		return ""
	}
	end := len(s) - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return ""
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	return s[start+1 : end+1]
}

func main() {
	fmt.Println(lastword("Hana ahmed gamal  "))

}
