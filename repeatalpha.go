package main

import "fmt"

func repeatalpha(s string) string {
	var result []rune
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			count = int((c - 'a') + 1)
		} else if c >= 'A' && c <= 'Z' {
			count = int((c - 'A') + 1)
		} else {
			count = 1
		}
		for i := 0; i < count; i++ {
			result = append(result, c)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(repeatalpha("Hana ahmed gamal"))
}
