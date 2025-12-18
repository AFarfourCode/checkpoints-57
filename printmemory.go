package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	for i := 0; i < len(arr); i += 4 {
		for j := 0; j < 4; j++ {
			if i+j < 10 {
				x := arr[i+j]
				z01.PrintRune(rune(hex[x%16]))
				z01.PrintRune(rune(hex[x/16]))
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
	for _, c := range arr {
		if c >= 32 && c <= 126 {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune('.')
		}
	}
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
