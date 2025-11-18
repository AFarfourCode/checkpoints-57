package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2 {
		return
	}
	for i := 2; i*i <= n; i++ {

		for n%i == 0 {
			fmt.Print(i)
			n /= i
			if n > 1 {
				fmt.Print("*")
			}
		}
	}
	if n > 1 {
		fmt.Print(n)
	}

	fmt.Println()
}
