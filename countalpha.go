<<<<<<< HEAD
package main

import "fmt"

func CountAlpha(s string) int {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			count++
		}
	}
	return count
}

func main() {
	x := "Hana ahmed"
	fmt.Println(CountAlpha(x))
}
=======
package main

import "fmt"

func CountAlpha(s string) int {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			count++
		}
	}
	return count
}

func main() {
	x := "Hana ahmed"
	fmt.Println(CountAlpha(x))
}
>>>>>>> 2262bbcac982957294399b14cf2d87d0af9024cf
