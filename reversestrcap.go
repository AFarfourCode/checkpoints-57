package main

import (
	"fmt"
	"os"
)

func reverseStrCap(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] != ' ' && (i == len(r)-1 || r[i+1] == ' ') {
			if r[i] >= 'a' && r[i] <= 'z' {
				r[i] -= 32
			}
		} else {
			if r[i] >= 'A' && r[i] <= 'Z' {
				r[i] += 32
			}
		}
	}
	return string(r)
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]

	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(reverseStrCap(arg))
	}
	fmt.Println()
}
