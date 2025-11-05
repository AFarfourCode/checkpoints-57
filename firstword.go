package main

import "fmt"

func firstWord(s string) string {
	result := ""
	flag := false
	for _, c := range s {

		if c != ' ' {
			result += string(c)
			flag = true
		}
		if c == ' ' && flag == true {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(firstWord("Hana Ahmed Gamal"))
	fmt.Println(firstWord("hello world"))
	fmt.Println(firstWord("hello   .........  bye"))
}
