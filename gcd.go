package main

import "fmt"

func GCD(a, b uint) uint {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(GCD(42, 10))
	fmt.Println(GCD(42, 12))
	fmt.Println(GCD(14, 77))
	fmt.Println(GCD(17, 3))
}
