package main

import "fmt"

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for n := nb; n >= 2; n-- {
		isPrime := true

		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return n
		}
	}

	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
