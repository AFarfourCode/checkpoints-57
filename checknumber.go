<<<<<<< HEAD
package main

import "fmt"

func CheckNumber(arg string) bool {

	for _, c := range arg {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func main() {
	x := "hello1"
	y := "hello"
	fmt.Println(CheckNumber(x))
	fmt.Println(CheckNumber(y))
}
=======
package main

import "fmt"

func CheckNumber(arg string) bool {

	for _, c := range arg {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}

func main() {
	x := "hello1"
	y := "hello"
	fmt.Println(CheckNumber(x))
	fmt.Println(CheckNumber(y))
}
>>>>>>> 2262bbcac982957294399b14cf2d87d0af9024cf
