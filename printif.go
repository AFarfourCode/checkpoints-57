<<<<<<< HEAD
package main

import "fmt"

func PrintIf(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) > 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
=======
package main

import "fmt"

func PrintIf(str string) string {
	if str == "" {
		return "G\n"
	}
	if len(str) > 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}
>>>>>>> 2262bbcac982957294399b14cf2d87d0af9024cf
