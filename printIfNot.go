<<<<<<< HEAD
package main

import "fmt"

func PrintIfNot(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
=======
package main

import "fmt"

func PrintIfNot(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIfNot("abcdefz"))
	fmt.Print(PrintIfNot("abc"))
	fmt.Print(PrintIfNot(""))
	fmt.Print(PrintIfNot("14"))
}
>>>>>>> 2262bbcac982957294399b14cf2d87d0af9024cf
