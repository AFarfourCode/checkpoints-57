package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	for i := 0; i < len(s); i++ {
		c := s[i]

		if !(c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z') {
			return s
		}
		if i > 0 && c >= 'A' && c <= 'Z' && c >= s[i-1] && c <= s[i-1] {
			return s
		}
		last := s[len(s)-1]
		if last >= 'A' && last <= 'Z' {
			return s
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {

		ch := s[i]
		if i != 0 {
			if ch >= 'A' && ch <= 'Z' {
				result += "_"
			}
		}
		result += string(ch)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
