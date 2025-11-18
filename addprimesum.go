/*
package main

import "fmt"

func addprimesum(n int) int {
	if n < 2 {
		return 0
	}

	result := 0

	for i := 2; i <= n; i++ {
		isprime := true

		// check if i is prime
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			result += i
		}
	}

	return result
}

func main() {
	fmt.Println(addprimesum(5))  // 10
	fmt.Println(addprimesum(7))  // 17
	fmt.Println(addprimesum(-2)) // 0
	fmt.Println(addprimesum(0))  // 0
}
*/

package main

import "fmt"

func addprimesum(n int) int {
	if n < 2 {
		return 0
	}
	result := 0
	for i := 2; i <= n; i++ {
		isprime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			result += i
		}
	}
	return result
}

func main() {
	fmt.Println(addprimesum(5))  // 10
	fmt.Println(addprimesum(7))  // 17
	fmt.Println(addprimesum(-2)) // 0
	fmt.Println(addprimesum(0))  // 0
}
