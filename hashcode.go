package main

import "fmt"

func HashCode(dec string) string {
	result := ""
	length := len(dec)

	for _, c := range dec {
		hash := (int(c) + length) % 127
		if hash < 33 {
			hash += 33
		}
		result += string(rune(hash))
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
