package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) < 3 {
		return ""
	}
	result := ""
	for i := 2; i < len(str); i += 3 {
		ch := str[i]
		result += string(ch)
	}
	return result
}

func main() {
	fmt.Println(ThirdTimeIsACharm("123456789"))
	fmt.Println(ThirdTimeIsACharm(""))
	fmt.Println(ThirdTimeIsACharm("a b c d e f"))
	fmt.Println(ThirdTimeIsACharm("12"))
}
